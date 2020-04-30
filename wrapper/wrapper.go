package wrapper

import (
	"github.com/button-tech/logger"
	"github.com/button-tech/utils-eth-tokens-getter/contract"
	"github.com/button-tech/utils-eth-tokens-getter/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"strconv"
	"sync"
)

type TokenInfo struct {
	Token    string  `json:"token"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Decimals int     `json:"decimals"`
}

type Token struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	TokenID  string `json:"token_id"`
	Coin     int    `json:"coin"`
}

func GetTokensBalancesByAddress(user common.Address, endpoint string, tokens []Token, result chan []TokenInfo) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		logger.Error(err)
		return
	}

	instance, err := contract.NewContract(common.HexToAddress(storage.ContractAddress), client)
	if err != nil {
		logger.Error(err)
		return
	}

	resultChan := make(chan TokenInfo, len(tokens))
	var wg sync.WaitGroup
	wg.Add(len(tokens))
	for _, token := range tokens {
		go func(token Token, wg *sync.WaitGroup) {
			defer wg.Done()
			tokenBalance, err := instance.GetTokenBalance(&bind.CallOpts{}, user, common.HexToAddress(token.TokenID))
			if err != nil {
				logger.Error("ParseFloat", err.Error())
				return
			}
			balanceStr := tokenBalance.String()
			if balanceStr != "0" {
				balanceFloat, err := strconv.ParseFloat(balanceStr, 64)
				if err != nil {
					logger.Error("ParseFloat", err.Error())
					return
				}

				decFloat := float64(token.Decimals)

				resultFloat := balanceFloat / math.Pow(10, decFloat)

				resultChan <- TokenInfo{Amount: resultFloat, Currency: token.Symbol, Token: token.TokenID, Decimals: token.Decimals}
			}
		}(token, &wg)
	}
	wg.Wait()
	close(resultChan)

	tokenBalances := make([]TokenInfo, 0, len(resultChan))
	for r := range resultChan {
		tokenBalances = append(tokenBalances, r)
	}
	result <- tokenBalances
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
