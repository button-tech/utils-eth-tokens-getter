package storage

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

func (s *StoredEthEndpoints) AddEthEndpoints(entry schema.EthEntry) {
	s.Lock()
	s.EthEndpoints = entry
	s.Unlock()
}

func (s *StoredEthEndpoints) GetEthEndpoints() *schema.EthEntry {
	s.RLock()
	defer s.RUnlock()
	return &s.EthEndpoints
}


func StartEthEndpointsStoring() {

	log.Println("Start ETH endpoints storing!")

	for {
		entry, err := db.GetEthEntries()
		if err != nil {
			log.Println(err)
			continue
		}

		EthEndpointsFromStorage.AddEthEndpoints(*entry)

		time.Sleep(time.Minute * 10)
	}
}


func GetEthEndpoints() ([]string, error) {
	endpoints := EthEndpointsFromStorage.GetEthEndpoints()
	if endpoints == nil {
		return nil, errors.New("Not found")
	}

	return endpoints.Addresses, nil
}
