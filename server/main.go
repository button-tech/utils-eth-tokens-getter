package main

import (
	"github.com/button-tech/utils-eth-tokens-getter/contract_monitoring"
	"github.com/button-tech/utils-eth-tokens-getter/server/handlers"
	"github.com/button-tech/utils-eth-tokens-getter/storage"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func init() {

	go storage.StartEthEndpointsStoring()
	go storage.StartTokenListStoring()
	go contract_monitoring.StartMonitoring()

	storage.ContractAddress = os.Getenv("ADDRESS")

	log.Println(storage.ContractAddress)
}

func main() {
	r := routing.New()

	r.Get("/tokenBalance/<address>", handlers.LookForTokens)

	if err := fasthttp.ListenAndServe(":8080", r.HandleRequest); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
