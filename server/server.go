package server

import (
	"github.com/button-tech/utils-eth-tokens-getter/contract-wrapper"
	"github.com/button-tech/utils-eth-tokens-getter/estorage"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"time"
)

type UserTokenBalances struct {
	Data []TokenInfo `json:"data"`
}

type TokenInfo struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Token    string `json:"token"`
}

func LookForTokens(c *gin.Context) {

	var (
		tokenAddresses []string
		tokenSymbols   []string
		tokenDec       []string
		balance        UserTokenBalances
		userAddress    = c.Param("address")
		result         = make(chan []string)
		errChan        = make(chan error)
	)

	for _, j := range singleton.TokenList {
		tokenAddresses = append(tokenAddresses, j.Address)
		tokenSymbols = append(tokenSymbols, j.Symbol)
		tokenDec = append(tokenDec, j.Dec)
	}

	endpoints, err := estorage.GetEthEndpoints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, e := range endpoints {
		go contract_wrapper.RequestTokenBalance(common.HexToAddress(userAddress), e, tokenAddresses, result, errChan)
	}

	select {
	case contractAnswer := <-result:
		for i := 0; i < len(contractAnswer); i++ {
			if contractAnswer[i] != "0" {

				balanceFloat, err := strconv.ParseFloat(contractAnswer[i], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err)
					return
				}

				decFloat, err := strconv.ParseFloat(tokenDec[i], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err)
					return
				}

				resultFloat := balanceFloat / math.Pow(10, decFloat)

				resultStr := strconv.FormatFloat(resultFloat, 'f', 8, 64)

				balance.Data = append(balance.Data, TokenInfo{Amount: resultStr, Currency: tokenSymbols[i], Token: tokenAddresses[i]})
			}
		}

		if len(balance.Data) == 0 {
			balance.Data = []TokenInfo{}
		}

		c.JSON(http.StatusOK, balance)
	case err := <-errChan:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	case <-time.After(2 * time.Second):
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bad request or timeout"})
	}
}
