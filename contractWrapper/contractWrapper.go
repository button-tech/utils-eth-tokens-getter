package contractWrapper

import (
	"github.com/button-tech/utils-eth-tokens-getter/contract"
	"github.com/button-tech/utils-eth-tokens-getter/estorage"
	"github.com/button-tech/utils-eth-tokens-getter/singleton"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func RequestBalancesForUsersOnContract(users []string, tokens []string) ([]string, error) {
	endpoint, err := estorage.GetEthEndpoint()
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	instance, err := balanceContract.NewBalanceContract(common.HexToAddress(singleton.ContractAddress), client)
	if err != nil {
		return nil, err
	}
	response, err := instance.GetTokenBalance(&bind.CallOpts{}, FromStringToCommonAddress(users), FromStringToCommonAddress(tokens))
	if err != nil {
		return nil, err
	} else {
		return FromBigIntToString(response), nil
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
