package server

import (
	"github.com/button-tech/utils-eth-tokens-getter/contract-wrapper"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/button-tech/utils-eth-tokens-getter/estorage"
	"time"
	"errors"
	"strconv"
	"math"
)

type UserTokenBalances struct {
	Data []TokenInfo `json:"data"`
}

type TokenInfo struct {
	Currency  string  `json:"currency"`
	Amount string `json:"amount"`
	Token string `json:"token"`
}

func LookForTokens(c *gin.Context) {

	var (
		tokenAddresses []string
	    tokenSymbols []string
		tokenDec []string
		balance UserTokenBalances
		userAddress = c.Param("address")
		result = make(chan []string)
	)

	for _, j := range singleton.TokenList {
		tokenAddresses = append(tokenAddresses, j.Address)
		tokenSymbols = append(tokenSymbols, j.Symbol)
		tokenDec = append(tokenDec, j.Dec)
	}

	endpoints, err := estorage.GetEthEndpoints()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, e := range endpoints{
		go contract_wrapper.RequestTokenBalance(common.HexToAddress(userAddress), e, tokenAddresses, result)
	}

	select {
	case contractAnswer := <- result:
		for i := 0; i < len(contractAnswer); i++ {
			if contractAnswer[i] != "0" {

				balanceFloat, err := strconv.ParseFloat(contractAnswer[i], 64 )
				if err != nil{
					c.JSON(http.StatusInternalServerError, err)
					return
				}

				decFloat, err := strconv.ParseFloat(tokenDec[i], 64)
				if err != nil{
					c.JSON(http.StatusInternalServerError, err)
					return
				}
				
				resultFloat := balanceFloat / math.Pow(10, decFloat)

				resultStr := strconv.FormatFloat(resultFloat, 'f', 8 , 64)

				balance.Data = append(balance.Data, TokenInfo{Amount: resultStr, Currency: tokenSymbols[i], Token: tokenAddresses[i]})
			}
		}
		c.JSON(http.StatusOK, balance)
	case <-time.After(2 * time.Second):
		c.JSON(http.StatusInternalServerError, errors.New("Bad request or timeout"))
	}

}
