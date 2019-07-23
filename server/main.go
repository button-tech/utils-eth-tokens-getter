package main

import (
	"github.com/button-tech/utils-eth-tokens-getter/server/handlers"
	"github.com/button-tech/utils-eth-tokens-getter/storage"
	"github.com/button-tech/utils-eth-tokens-getter/storage/endpoints"
	"github.com/button-tech/utils-eth-tokens-getter/token-list"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func init() {

	go endpoints.StartStoring()

	tokens, err := token_list.GetTokenList("./token-list/tokenList.csv")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	storage.ContractAddress = os.Getenv("ADDRESS")
	storage.TokenList = tokens

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
