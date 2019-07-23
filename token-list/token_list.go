package token_list

import (
	"encoding/csv"
	"os"
)

type Tokens struct {
	Symbol  string
	Address string
	Dec     string
}

func GetTokenList(file string) ([]Tokens, error) {
	f, err := os.Open(file)
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
			Dec:     line[2],
		}
		data = append(data, token)
	}

	return data, nil
}
