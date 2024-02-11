package ardor

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func (a *Ardor) ArdorGetAssetsByIssuerAccount(accountRS string, network string, start int, end int) (response ArdorResponse, err error) {
	url := fmt.Sprintf("%s?requestType=getAssetsByIssuer&account=%s&firstIndex=%d&lastIndex=%d", a.Endpoint, accountRS, start, end)

	res, err := a.GetRequest(url)
	if err != nil {
		return ArdorResponse{}, err
	}

	if res.ErrorCode != 0 {
		return ArdorResponse{}, fmt.Errorf("%s", res.ErrorDescription)
	}

	return res, nil
}

func (a *Ardor) ArdorGetAccountAssets(accountRS string, network string) (response ArdorResponse, err error) {
	url := fmt.Sprintf("%s?requestType=getAccountAssets&account=%s&includeAssetInfo=true", a.Endpoint, accountRS)
	resArdor, errArdor := a.GetRequest(url)
	if errArdor != nil {
		return ArdorResponse{}, errArdor
	}

	if resArdor.ErrorCode != 0 {
		err = fmt.Errorf("%s", resArdor.ErrorDescription)
		return ArdorResponse{}, err
	}
	return resArdor, nil
}

// ardorGetAccountAssets - get the list of user's assets
func (a *Ardor) ArdorGetAccountOneAsset(accountRS string, network string, asset string) (response ArdorResponse, err error) {
	url := fmt.Sprintf("%s?requestType=getAccountAssets&account=%s&includeAssetInfo=true&asset=%s", a.Endpoint, accountRS, asset)
	resArdor, errArdor := a.GetRequest(url)
	if errArdor != nil {
		return ArdorResponse{}, errArdor
	}

	if resArdor.ErrorCode != 0 {
		err = fmt.Errorf("%s", resArdor.ErrorDescription)
		return ArdorResponse{}, err
	}
	return resArdor, nil
}

// transferAsset - transfer asset to another account
func (a *Ardor) TransferAsset(network string, recipient string, senderPublicKey string, asset string, encryptedMessage string, quantityQNT int) (res ArdorJsonResponse, err error) {
	_url := a.buildURL("?requestType=transferAsset")

	data := url.Values{}
	data.Set("chain", "2")
	data.Set("recipient", recipient)
	data.Set("publicKey", senderPublicKey)
	data.Set("asset", asset)
	data.Set("quantityQNT", strconv.Itoa(quantityQNT))
	data.Set("deadline", "100")
	data.Set("feeNQT", "-1")
	data.Set("broadcast", "false")
	data.Set("message", encryptedMessage)
	// data.Set("messageToEncrypt", encryptedMessage)
	// data.Set("messageToEncryptIsText", "true")

	res, err = a.PostUrlencodedRequest(_url, data, 5)
	if err != nil {
		return
	}
	_, ok := res["errorCode"]
	if ok {
		err = fmt.Errorf("ardor error: %s", res["errorDescription"])
		return
	}

	return
}

func (a *Ardor) prepareUnsignedSendAsset(network string, accountRS string, recipientPublicKey string, assetID string, quantity uint64) error {
	senderPublicKey := os.Getenv("SIGBRO_RS1_PUBLIC_KEY")

	// let's get blockchain status to calculate finish height
	blockchainStatus, errBlockchain := a.getBlockchainStatus(network)
	if errBlockchain != nil {
		return errBlockchain
	}

	params := map[string]interface{}{
		"requestType":                 "transferAsset",
		"chain":                       2,
		"recipient":                   accountRS,
		"asset":                       assetID,
		"quantityQNT":                 quantity,
		"publicKey":                   senderPublicKey,
		"feeNQT":                      -1,
		"deadline":                    360,
		"broadcast":                   false,
		"recipientPublicKey":          recipientPublicKey,
		"phased":                      true,
		"phasingVotingModel":          7,
		"phasingQuorum":               1,
		"phasingMinBalance":           0,
		"phasingFinishHeight":         blockchainStatus.NumberOfBlocks + 100,
		"phasingSenderPropertySetter": "12347953972509949998",
		"phasingSenderPropertyName":   "iremember",
		"phasingSenderPropertyValue":  "whitelistedAccount",
	}

	payload := encodeParams(params)

	response, err := a.PostRawRequest(a.Endpoint, payload, 10)
	if err != nil {
		return err
	}

	var jsonResponse ArdorResponse

	// parse response
	jsonErr := json.Unmarshal(response, &jsonResponse)
	if jsonErr != nil {
		return jsonErr
	}

	if jsonResponse.ErrorCode != 0 {
		return fmt.Errorf("%s", jsonResponse.ErrorDescription)
	}

	unsignedFn := os.Getenv("SIGBRO_JS_UNSIGNED_TX")
	f, fErr := os.Create(unsignedFn)
	if fErr != nil {
		return fErr
	}
	defer f.Close()

	transactionJson, _ := json.Marshal(jsonResponse.TransactionJSON)
	_, wErr := f.Write([]byte(string(transactionJson)))
	if wErr != nil {
		return wErr
	}

	_ = f.Close()

	return nil
}

func (a *Ardor) getAllAssetsByIssuer(account string, network string) (items []ArdorAsset) {
	step := 100
	start := 0
	end := 100

	for {
		assets, assertsErr := a.ArdorGetAssetsByIssuerAccount(account, network, start, end)
		if assertsErr != nil {
			return
		}

		if assets.Assets != nil && len(assets.Assets) > 0 {
			assetsList := assets.Assets[0]
			count := 0
			for _, item := range assetsList {
				count += 1
				items = append(items, item)
			}

			if count < step {
				break
			}
		} else {
			break
		}

		start += step
		end += step
	}

	return
}
