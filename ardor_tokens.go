package ardor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"io"
)

func ardorDecodeToken(ardorEndpoint string, key string, token string) (response ArdorResponse, err error) {

	timeout, _ := strconv.Atoi(httpTimeout)
	httpClient := http.Client{Timeout: time.Second * time.Duration(timeout)}

	ardorResp := ArdorResponse{}

	url := fmt.Sprintf("%s?requestType=decodeToken&website=%s&token=%s", ardorEndpoint, key, token)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ardorResp, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", appName, appVersion))

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return ardorResp, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(io.Reader(res.Body))
	if readErr != nil {
		return ardorResp, readErr
	}

	jsonErr := json.Unmarshal(body, &ardorResp)
	if readErr != nil {
		return ardorResp, jsonErr
	}
	return ardorResp, nil
}
