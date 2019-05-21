package main

import (
	"github.com/EnoRage/ethbal/server"
	"github.com/EnoRage/ethbal/singleton"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateEthereumClient() *ethclient.Client {
	client, err := ethclient.Dial("")
	if err != nil {
		log.Println(err)
	}
	return client
}

func CreateGinServer() *gin.Engine {
	ginServer := gin.New()
	return ginServer
}

func RunGinServer(ginServer *gin.Engine) {
	ginServer.Use(gin.Recovery())
	ginServer.Use(gin.Logger())
	ginServer.POST("/balances", server.LookForTokens)
	ginServer.Run(":80")
}

func init() {
	// check for ping
	singleton.Client = CreateEthereumClient()
	singleton.GinServer = CreateGinServer()
}

func main() {
	RunGinServer(singleton.GinServer)
}
