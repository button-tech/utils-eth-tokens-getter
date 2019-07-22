package server

import (
	"github.com/button-tech/utils-eth-tokens-getter/contract-wrapper"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	userAddress := c.Param("address")

	var tokenAddresses []string
	var tokenSymbols []string

	for _, j := range singleton.TokenList {
		tokenAddresses = append(tokenAddresses, j.Address)
		tokenSymbols = append(tokenSymbols, j.Symbol)
	}

	contractAnswer, err := contract_wrapper.RequestBalancesForUsersOnContract(common.HexToAddress(userAddress), tokenAddresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(contractAnswer)

	var balance UserBalance

	for i := 0; i < len(contractAnswer); i++ {
		if contractAnswer[i] != "0" {
			balance.TokenBalanceGroup = append(balance.TokenBalanceGroup, TokenAndBalance{Balance: contractAnswer[i], Symbol: tokenSymbols[i]})
		}
	}

	c.JSON(http.StatusOK, balance)
}
