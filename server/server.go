package server

import (
	"github.com/button-tech/utils-eth-tokens-getter/contractWrapper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersAndToken struct {
	Users  []string `json:"users"`
	Tokens []string `json:"tokens"`
}

type UserBalance struct {
	User              string            `json:"user"`
	TokenBalanceGroup []TokenAndBalance `json:"token_balance_group"`
}

type TokenAndBalance struct {
	TokenAddress string `json:"token_address"`
	Balance      string `json:"balance"`
}

func LookForTokens(c *gin.Context) {
	var reqData UsersAndToken
	err := c.BindJSON(&reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	contractAnswer, err := contractWrapper.RequestBalancesForUsersOnContract(reqData.Users, reqData.Tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	amountOfUsers := len(reqData.Users)
	amountOfTokens := len(reqData.Tokens)

	var userBalances = make([]UserBalance, 0)
	for i := 0; i < amountOfUsers; i++ {
		var userBalance UserBalance
		var tokenBalanceGroup = make([]TokenAndBalance, 0)
		var lhs = amountOfTokens * i
		var rhs = lhs + amountOfTokens
		for j := lhs; j < rhs; j++ {
			var tokenBalance TokenAndBalance
			tokenBalance.Balance = contractAnswer[j]
			tokenBalance.TokenAddress = reqData.Tokens[j-lhs]
			tokenBalanceGroup = append(tokenBalanceGroup, tokenBalance)
		}
		userBalance.User = reqData.Users[i]
		userBalance.TokenBalanceGroup = tokenBalanceGroup
		userBalances = append(userBalances, userBalance)
	}

	c.JSON(http.StatusOK, userBalances)
}
