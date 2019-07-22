package estorage

import (
	"sync"
	"github.com/button-tech/utils-eth-tokens-getter/estorage/db/schema"
	"math/rand"
	"errors"
	"time"
	"log"
	"github.com/button-tech/utils-eth-tokens-getter/estorage/db"
)

type StoredEthEndpoints struct {
	sync.RWMutex
	EthEndpoints schema.EthEntry
}

var EthEndpointsFromStorage StoredEthEndpoints

func (s *StoredEthEndpoints) Add(entry schema.EthEntry) {
	s.Lock()
	s.EthEndpoints = entry
	s.Unlock()
}

func(s *StoredEthEndpoints) Get() *schema.EthEntry{
	s.RLock()
	defer s.RUnlock()
	return &s.EthEndpoints
}

func StartStoring(){

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

func GetEthEndpoint() (string, error){
	endpoints := EthEndpointsFromStorage.Get()
	if endpoints == nil {
		return "", errors.New("Not found")
	}

	addresses := endpoints.Addresses

	rand.Seed(time.Now().UnixNano())

	result := addresses[rand.Intn(len(addresses))]

	return result, nil
}