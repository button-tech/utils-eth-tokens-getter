package server

import (
	"github.com/EnoRage/ethbal/contractWrapper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersAndToken struct {
	Users  []string `json:"users"`
	Tokens []string `json:"tokens"`
}

type UserBalance struct {
	User              string `json:"user"`
	TokenBalanceGroup []TokenAndBalance
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
	contractAnswerLen := len(contractAnswer)
	amountOfUsers := len(reqData.Users)
	amountOfTokens := len(reqData.Tokens)
	upBorder := contractAnswerLen/ amountOfUsers

	var userBalances = make([]UserBalance, 0)
	for i := 0; i <= amountOfUsers; i++ {
		var userBalance UserBalance
		var tokenBalanceGroup = make([]TokenAndBalance, 0)
		for j := 0; j <= amountOfTokens; j++ {
			var tokenBalance TokenAndBalance
			tokenBalance.Balance = contractAnswer[]
			tokenBalance.TokenAddress = reqData.Tokens[i]
			tokenBalanceGroup = append(tokenBalanceGroup, tokenBalance)
		}
		userBalance.User = reqData.Users[i]
		userBalance.TokenBalanceGroup = tokenBalanceGroup
		userBalances = append(userBalances, userBalance)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, userBalances)
		return
	}
}
