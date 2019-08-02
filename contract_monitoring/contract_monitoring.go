package contract_monitoring

import (
	"github.com/button-tech/utils-eth-tokens-getter/contract_wrapper"
	"github.com/button-tech/utils-eth-tokens-getter/db"
	"github.com/button-tech/utils-eth-tokens-getter/storage"
	"log"
	"os"
	"time"
)

func StartMonitoring() {

	address := os.Getenv("ETH_ADDRESS_FOR_CHECK")

	log.Println("Start contract monitoring!")

	for {

		tokens := storage.TokenList.GetTokens()

		if len(tokens) == 0 {
			continue
		}

		for _, j := range tokens {

			err := contract_wrapper.GetTokenBalance(address, j.Address)
			if err != nil {
				if err.Error() == "abi: unmarshalling empty output" || err.Error() == "VM execution error." {
					err := db.DeleteToken(j.Address)
					if err != nil {
						log.Println("Can't delete:" + j.Address)
						continue
					} else {
						log.Println("Successfully delete:" + j.Address)
					}
				}
			}

			time.Sleep(300 * time.Millisecond)
		}

		log.Println("All token addresses was checked!")

	}
}
