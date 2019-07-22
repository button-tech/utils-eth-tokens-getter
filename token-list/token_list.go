package token_list

import (
	"encoding/csv"
	"os"
)

type Tokens struct {
	Symbol  string
	Address string
}

func GetTokenList() ([]Tokens, error) {
	f, err := os.Open("tokenList.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	var data []Tokens

	for _, line := range lines {
		token := Tokens{
			Symbol:  line[0],
			Address: line[1],
		}
		data = append(data, token)
	}

	return data, nil
}
