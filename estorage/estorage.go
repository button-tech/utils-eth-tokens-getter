package estorage

import (
	"errors"
	"github.com/button-tech/utils-eth-tokens-getter/db"
	"github.com/button-tech/utils-eth-tokens-getter/db/schema"
	"log"
	"sync"
	"time"
)

type StoredEthEndpoints struct {
	sync.RWMutex
	EthEndpoints schema.EthEntry
}

func (s *StoredEthEndpoints) Add(entry schema.EthEntry) {
	s.Lock()
	s.EthEndpoints = entry
	s.Unlock()
}

func (s *StoredEthEndpoints) Get() *schema.EthEntry {
	s.RLock()
	defer s.RUnlock()
	return &s.EthEndpoints
}

var EthEndpointsFromStorage StoredEthEndpoints

func StartStoring() {

	log.Println("Start storing!")

	for {
		entry, err := db.GetEthEntries()
		if err != nil {
			log.Println(err)
			continue
		}

		EthEndpointsFromStorage.Add(*entry)

		time.Sleep(time.Minute * 10)
	}
}

func GetEthEndpoints() ([]string, error) {
	endpoints := EthEndpointsFromStorage.Get()
	if endpoints == nil {
		return nil, errors.New("Not found")
	}

	return endpoints.Addresses, nil
}
