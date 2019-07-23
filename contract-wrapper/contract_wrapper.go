package contract_wrapper

import (
	"github.com/button-tech/utils-eth-tokens-getter/contract"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func RequestTokenBalance(user common.Address, endpoint string, tokens []string, result chan []string, errChan chan error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		errChan <- err
		return
	}

	instance, err := contract.NewContract(common.HexToAddress(singleton.ContractAddress), client)
	if err != nil {
		errChan <- err
		return
	}
	response, err := instance.GetTokenBalanceByAddress(&bind.CallOpts{}, user, FromStringToCommonAddress(tokens))
	if err != nil {
		errChan <- err
		return
	} else {
		result <- FromBigIntToString(response)
	}
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
