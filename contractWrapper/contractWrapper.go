package contractWrapper

import (
	"github.com/EnoRage/ethbal/contract"
	"github.com/EnoRage/ethbal/singleton"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func RequestBalancesForUsersOnContract(users []string, tokens []string) ([]string, error) {
	instance, err := balanceContract.NewBalanceContract(common.HexToAddress(singleton.ContractAddress), singleton.Client)
	if err != nil {
		return nil, err
	}
	response, err := instance.GetTokenBalance(&bind.CallOpts{}, FromStringToCommonAddress(users), FromStringToCommonAddress(tokens))
	if err != nil {
		return nil, err
	}
	return FromBigIntToString(response), nil
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
