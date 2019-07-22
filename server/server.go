package server

import (
	"github.com/button-tech/utils-eth-tokens-getter/contract-wrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
)


type UserBalance struct {
	TokenBalanceGroup []TokenAndBalance `json:"token_balance_group"`
}

type TokenAndBalance struct {
	TokenAddress string `json:"token_address"`
	Balance      string `json:"balance"`
}

func LookForTokens(c *gin.Context) {

	userAddress := c.Param("address")

	var tokenAddresses []string

	for _, j := range singleton.TokenList{
		tokenAddresses = append(tokenAddresses, j.Address)
	}

	contractAnswer, err := contract_wrapper.RequestBalancesForUsersOnContract(common.HexToAddress(userAddress), tokenAddresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	amountOfTokens := len(tokenAddresses)

	var userBalances = make([]UserBalance, 0)
	for i := 0; i < amountOfTokens; i++ {
		var userBalance UserBalance
		var tokenBalanceGroup = make([]TokenAndBalance, 0)
		var lhs = amountOfTokens * i
		var rhs = lhs + amountOfTokens
		for j := lhs; j < rhs; j++ {
			var tokenBalance TokenAndBalance
			tokenBalance.Balance = contractAnswer[j]
			tokenBalance.TokenAddress = tokenAddresses[j-lhs]
			tokenBalanceGroup = append(tokenBalanceGroup, tokenBalance)
		}

		userBalance.TokenBalanceGroup = tokenBalanceGroup
		userBalances = append(userBalances, userBalance)
	}

	c.JSON(http.StatusOK, userBalances)
}
