package main

import (
	"github.com/button-tech/utils-eth-tokens-getter/estorage"
	"github.com/button-tech/utils-eth-tokens-getter/server"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
	"github.com/button-tech/utils-eth-tokens-getter/token-list"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func CreateGinServer() *gin.Engine {
	ginServer := gin.New()
	return ginServer
}

func RunGinServer(ginServer *gin.Engine) {

	ginServer.Use(gin.Recovery())

	ginServer.GET("/tokenBalance/:address", server.LookForTokens)

	gin.SetMode(gin.ReleaseMode)

	err := ginServer.Run(":" + singleton.GinPort)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {

	go estorage.StartStoring()

	tokens, err := token_list.GetTokenList("./token-list/tokenList.csv")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	singleton.ContractAddress = os.Getenv("ADDRESS")
	singleton.GinPort = os.Getenv("GIN_PORT")
	singleton.TokenList = tokens

	log.Println(os.Getenv("ADDRESS"))
	log.Println(os.Getenv("GIN_PORT"))

	singleton.GinServer = CreateGinServer()
}

func main() {
	RunGinServer(singleton.GinServer)
}
