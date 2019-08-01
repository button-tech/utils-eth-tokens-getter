package handlers

import (
	"encoding/json"
	"errors"
	"github.com/button-tech/utils-eth-tokens-getter/contract_wrapper"
	"github.com/button-tech/utils-eth-tokens-getter/storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/qiangxue/fasthttp-routing"
	"math"
	"strconv"
	"time"
)

type UserTokenBalances struct {
	Data []TokenInfo `json:"data"`
}

type TokenInfo struct {
	Token    string  `json:"token"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func LookForTokens(c *routing.Context) error {

	var (
		tokenAddresses []string
		tokenSymbols   []string
		tokenDec       []string
		balance        UserTokenBalances
		userAddress    = c.Param("address")
		result         = make(chan []string)
		errChan        = make(chan error)
	)

	tokenList := storage.TokenList.GetTokens()

	for _, j := range tokenList {
		tokenAddresses = append(tokenAddresses, j.Address)
		tokenSymbols = append(tokenSymbols, j.Symbol)
		tokenDec = append(tokenDec, j.Decimal)
	}

	es, err := storage.GetEthEndpoints()
	if err != nil {
		return err
	}

	for _, e := range es {
		go contract_wrapper.GetTokensBalancesByAddress(common.HexToAddress(userAddress), e, tokenAddresses, result, errChan)
	}

	select {
	case contractAnswer := <-result:
		for i := 0; i < len(contractAnswer); i++ {
			if contractAnswer[i] != "0" {

				balanceFloat, err := strconv.ParseFloat(contractAnswer[i], 64)
				if err != nil {
					return err
				}

				decFloat, err := strconv.ParseFloat(tokenDec[i], 64)
				if err != nil {
					return err
				}

				resultFloat := balanceFloat / math.Pow(10, decFloat)

				balance.Data = append(balance.Data, TokenInfo{Amount: resultFloat, Currency: tokenSymbols[i], Token: tokenAddresses[i]})
			}
		}

		if len(balance.Data) == 0 {
			balance.Data = []TokenInfo{}
		}

		err := JsonResponse(c, balance)
		if err != nil {
			return err
		}

	case err := <-errChan:
		return err
	case <-time.After(2 * time.Second):
		return errors.New("Bad request or timeout")
	}

	return nil
}

func JsonResponse(ctx *routing.Context, data interface{}) error {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "https://client.buttonwallet.tech")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, HEAD")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	ctx.Response.SetStatusCode(200)
	if err := json.NewEncoder(ctx).Encode(data); err != nil {
		return err
	}
	return nil
}
