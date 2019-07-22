package server

import (
	"errors"
	"github.com/button-tech/utils-eth-tokens-getter/contract-wrapper"
	"github.com/button-tech/utils-eth-tokens-getter/estorage"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserBalance struct {
	TokenBalanceGroup []TokenAndBalance `json:"token_balance_group"`
}

type TokenAndBalance struct {
	//TokenAddress string `json:"token_address"`
	Symbol  string
	Balance string `json:"balance"`
}

func LookForTokens(c *gin.Context) {

	var (
		tokenAddresses []string
		tokenSymbols   []string
		balance        UserBalance
		userAddress    = c.Param("address")
		result         = make(chan []string)
	)

	for _, j := range singleton.TokenList {
		tokenAddresses = append(tokenAddresses, j.Address)
		tokenSymbols = append(tokenSymbols, j.Symbol)
	}

	endpoints, err := estorage.GetEthEndpoints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, e := range endpoints {
		go contract_wrapper.RequestTokenBalance(common.HexToAddress(userAddress), e, tokenAddresses, result)
	}

	select {
	case contractAnswer := <-result:
		for i := 0; i < len(contractAnswer); i++ {
			if contractAnswer[i] != "0" {
				balance.TokenBalanceGroup = append(balance.TokenBalanceGroup, TokenAndBalance{Balance: contractAnswer[i], Symbol: tokenSymbols[i]})
			}
		}
		c.JSON(http.StatusOK, balance)
	case <-time.After(2 * time.Second):
		c.JSON(http.StatusInternalServerError, errors.New("Bad request or timeout"))
	}

}
