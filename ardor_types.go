package ardor

type Ardor struct {
	Endpoint string
}

type ArdorJsonResponse map[string]interface{}

type ArdorResponse struct {
	ErrorCode                int64                  `json:"errorCode"`
	ErrorDescription         string                 `json:"errorDescription"`
	AccountRS                string                 `json:"accountRS"`
	PublicKey                string                 `json:"publicKey"`
	FullHash                 string                 `json:"fullHash"`
	QuantityQNT              string                 `json:"quantityQNT"`
	TransactionJSON          map[string]interface{} `json:"transactionJSON"`
	UnsignedTransactionBytes string                 `json:"unsignedTransactionBytes"`
	Assets                   [][]ArdorAsset         `json:"assets"`
	AccountAssets            []ArdorAsset           `json:"accountAssets"` // getAssetAccounts
	Properties               []ArdorAssetProperty   `json:"properties"`
	SetterRS                 string                 `json:"setterRS"` // getAccountProperties
	Transaction              string                 `json:"transaction"`
	Trades                   []ArdorTrades          `json:"trades"`
	UnixTime                 int64                  `json:"unixtime"`
	Time                     int64                  `json:"time"`
}

type ArdorTrades struct {
	QuantityQNT     string `json:"quantityQNT"`
	Chain           int    `json:"chain"`
	OrderFullHash   string `json:"orderFullHash"`
	ExcangeRate     string `json:"exchangeRate"`
	AccountRS       string `json:"accountRS"`
	Excange         int    `json:"exchange"`
	Block           string `json:"block"`
	MatchFullHash   string `json:"matchFullHash"`
	PriceNQTPerCoin string `json:"priceNQTPerCoin"`
	Account         string `json:"account"`
	Height          int    `json:"height"`
	Timestamp       int    `json:"timestamp"`
}

type ArdorAsset struct {
	QuantityQNT            string `json:"quantityQNT"`
	NumberOfAccounts       int    `json:"numberOfAccounts"`
	AccountRS              string `json:"accountRS"`
	Decimals               int    `json:"decimals"`
	NumberOfTransfers      int    `json:"numberOfTransfers"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	HasPhasingAssetControl bool   `json:"hasPhasingAssetControl"`
	Account                string `json:"account"`
	Asset                  string `json:"asset"`
}

type ArdorAssetProperty struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

type ArdorBlockchainStatusResponse struct {
	BlockchainState string `json:"blockchainState"`
	NumberOfBlocks  int64  `json:"numberOfBlocks"`
}
