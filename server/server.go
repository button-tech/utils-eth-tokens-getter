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

type UsersAndTokenAndBalances struct {
	UsersAndToken
	Balances []string `json:"balances"`
}

func LookForTokens(c *gin.Context) {
	var reqData UsersAndToken
	var serverAnswer UsersAndTokenAndBalances
	c.BindJSON(&reqData)
	serverAnswer.UsersAndToken = reqData
	contractAnswer, err := contractWrapper.RequestBalancesForUsersOnContract(reqData.Users, reqData.Tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	serverAnswer.Balances = contractAnswer
	c.JSON(http.StatusOK, serverAnswer)
}
