package main

import (
	"fmt"
	"github.com/EnoRage/ethbal/server"
	"github.com/EnoRage/ethbal/singleton"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func CreateEthereumClient() *ethclient.Client {
	client, err := ethclient.Dial(singleton.EthEndpoint)
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
	gin.SetMode(gin.ReleaseMode)
	err := ginServer.Run(":" + singleton.GinPort)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	// check for ping
	singleton.EthEndpoint = os.Getenv("ETH_ENDPOINT")
	singleton.ContractAddress = os.Getenv("ADDRESS")
	singleton.GinPort = os.Getenv("GIN_PORT")
	singleton.Client = CreateEthereumClient()
	singleton.GinServer = CreateGinServer()

	fmt.Println(os.Getenv("ETH_ENDPOINT"))
	fmt.Println(os.Getenv("ADDRESS"))
	fmt.Println(os.Getenv("GIN_PORT"))
}

func main() {
	RunGinServer(singleton.GinServer)
}
