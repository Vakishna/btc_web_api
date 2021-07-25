package models

type Wallet struct {
	Id           string `json:"id"`
	Address      string `json:"address"`
	Chain        int    `json:"chain"`
	Index        int    `json:"index"`
	Coin         string `json:"coin"`
	LastNonce    int    `json:"lastNonce"`
	Wallet       string `json:"wallet"`
	CoinSpecific string `json:"-"`
	Label        string `json:"-"`
	AddressType  string `json:"addressType"`
}
