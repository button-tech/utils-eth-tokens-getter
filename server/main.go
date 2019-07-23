package main

import (
	"github.com/button-tech/utils-eth-tokens-getter/estorage"
	"github.com/button-tech/utils-eth-tokens-getter/server/handlers"
	"github.com/button-tech/utils-eth-tokens-getter/shared"
	"github.com/button-tech/utils-eth-tokens-getter/token-list"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func init() {

	go estorage.StartStoring()

	tokens, err := token_list.GetTokenList("./token-list/tokenList.csv")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	shared.ContractAddress = os.Getenv("ADDRESS")
	shared.TokenList = tokens

	log.Println(shared.ContractAddress)
}

func main() {
	r := routing.New()

	r.Get("/tokenBalance/<address>", handlers.LookForTokens)

	if err := fasthttp.ListenAndServe(":8080", r.HandleRequest); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
