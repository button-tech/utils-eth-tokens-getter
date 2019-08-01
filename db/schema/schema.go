package schema

type EthEntry struct {
	Addresses []string `json:"addresses"`
	Port      int      `json:"port"`
}

type Tokens struct {
	Symbol string `json:"symbol"`
	Address string `json:"address"`
	Decimal string `json:"decimal"`
}
