package ardor

import (
	"encoding/json"
	"fmt"
	"os"
)

func (a *Ardor) prepareUnsignedSendMoney(network string, accountRS string, amount uint64) error {
	senderPublicKey := os.Getenv("SIGBRO_RS1_PUBLIC_KEY")
	payload := fmt.Sprintf("requestType=sendMoney&chain=2&recipient=%s&amountNQT=%d&publicKey=%s&feeNQT=-1&deadline=360&broadcast=false", accountRS, amount, senderPublicKey)

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

func (a *Ardor) prepareUnsignedSetAccountProperty(network string, accountRS string, publicKey string, property string) error {
	senderPublicKey := os.Getenv("SIGBRO_RS1_PUBLIC_KEY")
	payload := fmt.Sprintf("requestType=setAccountProperty&chain=2&recipient=%s&property=%s&publicKey=%s&feeNQT=-1&deadline=360&broadcast=false&recipientPublicKey=%s", accountRS, property, senderPublicKey, publicKey)

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
