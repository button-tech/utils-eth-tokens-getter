package handlers

import (
	"encoding/json"
	"errors"
	"github.com/button-tech/logger"
	"github.com/button-tech/utils-eth-tokens-getter/storage"
	"github.com/button-tech/utils-eth-tokens-getter/wrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/imroc/req"
	"github.com/qiangxue/fasthttp-routing"
	"os"
	"time"
)

type UserTokenBalances struct {
	Data []wrapper.TokenInfo `json:"data"`
}

type ApiResponse []wrapper.Token

func LookForTokens(c *routing.Context) error {
	start := time.Now()
	var (
		balance     UserTokenBalances
		userAddress = c.Param("address")
		result      = make(chan []wrapper.TokenInfo)
	)

	tokenList, err := GetTokensListByAddress(userAddress)
	if err != nil {
		logger.Error("GetTokensListByAddress", err.Error(), logger.Params{
			"userAddress": userAddress,
		})
		return err
	}

	if len(tokenList) == 0 {
		balance.Data = []wrapper.TokenInfo{}
		err := JsonResponse(c, balance)
		if err != nil {
			return err
		}
		return nil
	}

	es, err := storage.GetEthEndpoints()
	if err != nil {
		logger.Error("GetEthEndpoints", err.Error())
		return err
	}

	for _, e := range es {
		go wrapper.GetTokensBalancesByAddress(common.HexToAddress(userAddress), e, tokenList, result)
	}

	select {
	case contractAnswer := <-result:
		err := JsonResponse(c, UserTokenBalances{Data: contractAnswer})
		if err != nil {
			return err
		}
	case <-time.After(3 * time.Second):
		logger.Error("Bad request or timeout")
		return errors.New("Bad request or timeout")
	}

	logger.LogRequest(time.Since(start), "ETH", "GetTokenBalance", true)

	return nil
}

func GetTokensListByAddress(address string) (ApiResponse, error) {
	var result ApiResponse

	res, err := req.Get(os.Getenv("TOKEN_API") + address)
	if err != nil {
		return nil, err
	}

	err = res.ToJSON(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
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
