package handlers

import (
	"encoding/json"
	"errors"
	"github.com/button-tech/utils-eth-tokens-getter/contract_wrapper"
	"github.com/button-tech/utils-eth-tokens-getter/storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/imroc/req"
	"github.com/qiangxue/fasthttp-routing"
	"math"
	"os"
	"strconv"
	"time"
)

type UserTokenBalances struct {
	Data []TokenInfo `json:"data"`
}

type Token struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	TokenID  string `json:"tokenID"`
	Coin     int    `json:"coin"`
}

type ApiResponse struct {
	Docs []Token `json:"docs"`
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
		tokenDec       []int
		balance        UserTokenBalances
		userAddress    = c.Param("address")
		result         = make(chan []string)
		errChan        = make(chan error)
	)

	tokenList, err := GetTokensListByAddress(userAddress)
	if err != nil {
		return err
	}

	if len(tokenList.Docs) == 0 {
		balance.Data = []TokenInfo{}
		err := JsonResponse(c, balance)
		if err != nil {
			return err
		}
		return nil
	}

	for _, j := range tokenList.Docs {
		tokenAddresses = append(tokenAddresses, j.TokenID)
		tokenSymbols = append(tokenSymbols, j.Symbol)
		tokenDec = append(tokenDec, j.Decimals)
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

				decFloat := float64(tokenDec[i])

				resultFloat := balanceFloat / math.Pow(10, decFloat)

				balance.Data = append(balance.Data, TokenInfo{Amount: resultFloat, Currency: tokenSymbols[i], Token: tokenAddresses[i]})
			}
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

func GetTokensListByAddress(address string) (*ApiResponse, error) {
	var result ApiResponse

	res, err := req.Get(os.Getenv("TOKEN_API") + address)
	if err != nil {
		return nil, err
	}

	err = res.ToJSON(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
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
