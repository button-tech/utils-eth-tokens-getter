package main

import (
	"github.com/button-tech/utils-eth-tokens-getter/server"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"github.com/button-tech/utils-eth-tokens-getter/estorage"
	"github.com/button-tech/utils-eth-tokens-getter/estorage/db"
)

//init
func CreateEthereumClient(endpoint string) *ethclient.Client {
	client, err := ethclient.Dial(endpoint)
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
	ginServer.POST("/balances", server.LookForTokens)
	gin.SetMode(gin.ReleaseMode)
	err := ginServer.Run(":" + singleton.GinPort)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	// for first store
	ethEndpoints, err := db.GetEthEntries()
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}

	estorage.EthEndpointsFromStorage.Add(*ethEndpoints)

	// start storing
	go estorage.StartStoring()


	// check for ping
	singleton.ContractAddress = os.Getenv("ADDRESS")
	singleton.GinPort = os.Getenv("GIN_PORT")
	log.Println(os.Getenv("ADDRESS"))
	log.Println(os.Getenv("GIN_PORT"))

	endpoint, err := estorage.GetEthEndpoint()
	if err != nil{
		os.Exit(1)
	}

	singleton.Client = CreateEthereumClient(endpoint)
	singleton.GinServer = CreateGinServer()
}

func main() {
	RunGinServer(singleton.GinServer)
}