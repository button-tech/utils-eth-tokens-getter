package storage

import (
	"github.com/button-tech/utils-eth-tokens-getter/db"
	"github.com/button-tech/utils-eth-tokens-getter/db/schema"
	"log"
	"sync"
	"time"
)

type StoredTokenList struct {
	sync.RWMutex
	Tokens []schema.Tokens
}

func (t *StoredTokenList) AddTokens(entries []schema.Tokens) {
	t.Lock()
	t.Tokens = entries
	t.Unlock()
}

func (t *StoredTokenList) GetTokens() []schema.Tokens {
	t.RLock()
	defer t.RUnlock()
	return t.Tokens
}

func StartTokenListStoring() {

	log.Println("Start token list storing!")

	for {
		entries, err := db.GetTokensEntries()
		if err != nil {

			continue
		}

		TokenList.AddTokens(entries)

		time.Sleep(time.Minute * 1)
	}
}
