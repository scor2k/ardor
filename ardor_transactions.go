package ardor

import (
	"encoding/json"
	"fmt"
	"os"
)

func (a *Ardor) ardorGetTransaction(chain int, fullHash string, network string) (response ArdorJsonResponse, err error) {
	url := fmt.Sprintf("%s?requestType=getTransaction&chain=%d&fullHash=%s", a.Endpoint, chain, fullHash)
	response, err = a.GetRequestUnparsed(url)

	if err != nil {
		return ArdorJsonResponse{}, err
	}

	_, ok := response["errorCode"]
	if ok {
		err = fmt.Errorf("%s", response["errorDescription"].(string))
		return ArdorJsonResponse{}, err
	}
	return response, nil
}

func (a *Ardor) broadcastTransaction(network string) error {
	signedFn := os.Getenv("SIGBRO_JS_SIGNED_TX")

	signedTx, err := os.ReadFile(signedFn)
	if err != nil {
		return err
	}

	payload := fmt.Sprintf("requestType=broadcastTransaction&transactionJSON=%s", string(signedTx))

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

	return nil
}
