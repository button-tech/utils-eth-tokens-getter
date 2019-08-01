package contract_wrapper

import (
	"github.com/button-tech/utils-eth-tokens-getter/contract"
	"github.com/button-tech/utils-eth-tokens-getter/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetTokensBalancesByAddress(user common.Address, endpoint string, tokens []string, result chan []string, errChan chan error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		errChan <- err
		return
	}

	instance, err := contract.NewContract(common.HexToAddress(storage.ContractAddress), client)
	if err != nil {
		errChan <- err
		return
	}
	response, err := instance.GetTokensBalancesByAddress(&bind.CallOpts{}, user, FromStringToCommonAddress(tokens))
	if err != nil {
		errChan <- err
		return
	} else {
		result <- FromBigIntToString(response)
	}
}

func GetTokenBalance(userAddress, tokenAddress string) error {
	client, _ := ethclient.Dial("https://mainnet.infura.io")

	instance, _ := contract.NewContract(common.HexToAddress(storage.ContractAddress), client)

	_, err := instance.GetTokenBalance(&bind.CallOpts{}, common.HexToAddress(userAddress), common.HexToAddress(tokenAddress))
	if err != nil {
		return err
	}

	return nil
}

func FromStringToCommonAddress(args []string) []common.Address {
	response := make([]common.Address, 0)
	for i := range args {
		response = append(response, common.HexToAddress(args[i]))
	}
	return response
}

func FromBigIntToString(args []*big.Int) []string {
	response := make([]string, 0)
	for i := range args {
		response = append(response, args[i].String())
	}
	return response
}
