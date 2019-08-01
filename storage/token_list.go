package storage

import (
	"sync"
	"github.com/button-tech/utils-eth-tokens-getter/db/schema"
	"github.com/button-tech/utils-eth-tokens-getter/db"
	"time"
	"log"
)

type StoredTokenList struct {
	sync.RWMutex
	Tokens []schema.Tokens
}

func(t *StoredTokenList) AddTokens(entries []schema.Tokens){
	t.Lock()
	t.Tokens = entries
	t.Unlock()
}

func(t *StoredTokenList) GetTokens() []schema.Tokens {
	t.RLock()
	defer t.RUnlock()
	return t.Tokens
}

func StartTokenListStoring(){

	log.Println("Start token list storing!")

	for {
		entries, err := db.GetTokensEntries()
		if err != nil {

			continue
		}

		TokenList.AddTokens(entries)

		log.Println(TokenList.GetTokens())

		time.Sleep(time.Minute * 1)
	}
}




