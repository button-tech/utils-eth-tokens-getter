package contract

import "math/big"

func RequestBalancesForUsersOnContract(users []string, tokens []string) ([]string, error) {
	//instance, err := store.NewStore(contractAddress, Client)
	//if err != nil {
	//	return nil, err
	//}
	//instance

}

func FromBigIntToString(args []big.Int) []string {
	answer := make([]string, 0)
	for i := range args {
		answer = append(answer, args[i].String())
	}
	return answer
}
