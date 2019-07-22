package main

import (
	"github.com/button-tech/utils-eth-tokens-getter/server"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"github.com/button-tech/utils-eth-tokens-getter/estorage"
)

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

	// start storing
	go estorage.StartStoring()

	// check for ping
	singleton.ContractAddress = os.Getenv("ADDRESS")
	singleton.GinPort = os.Getenv("GIN_PORT")
	log.Println(os.Getenv("ADDRESS"))
	log.Println(os.Getenv("GIN_PORT"))
	singleton.GinServer = CreateGinServer()
}

func main() {
	RunGinServer(singleton.GinServer)
}