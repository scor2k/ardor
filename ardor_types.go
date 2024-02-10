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

// prepare unsigned Tx
type ArdorRequest struct {
	Description       string `json:"description"`
	Tags              string `json:"tags"`
	TypeName          string `json:"typeName"`
	Channel           string `json:"channel"`
	Data              string `json:"data"`
	PublicKey         string `json:"publicKey"`
	Deadline          int    `json:"deadline"`
	Broadcast         bool   `json:"broadcast"`
	MessageIsPrunable bool   `json:"messageIsPrunable"`
	QuantityQNT       int    `json:"quantityQNT"`
	Decimals          int    `json:"decimals"`
	Name              string `json:"name"`
	Chain             int    `json:"chain"`
	FeeNQT            string `json:"feeNQT"`
	AmountNQT         string `json:"amountNQT"`
	Type              int    `json:"type"`
	SubType           int    `json:"subtype"`
	SenderRS          string `json:"senderRS"`
	RecipientRS       string `json:"recipientRS"`
	FullHash          string `json:"fullHash"`
	Block             string `json:"block"`
	BlockHeight       uint64 `json:"height"`
	BlockTimestamp    uint64 `json:"blockTimestamp"`
	Timestamp         uint64 `json:"timestamp"`
	Transaction       string `json:"transaction"`
}

type ArdorBlockchainStatusResponse struct {
	BlockchainState string `json:"blockchainState"`
	NumberOfBlocks  int64  `json:"numberOfBlocks"`
}
