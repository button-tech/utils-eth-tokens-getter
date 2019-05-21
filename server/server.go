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

type Respond struct {
	UsersBalances []UserBalance `json:"users_balances"`
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
	var respond Respond
	err := c.BindJSON(&reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	contractAnswer, err := contractWrapper.RequestBalancesForUsersOnContract(reqData.Users, reqData.Tokens)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		// сюда все присвоить
		c.JSON(http.StatusOK, respond)
		return
	}
}
