package ardor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func encodeParams(params map[string]interface{}) string {
	var encodedParams []string

	for key, value := range params {
		var encodedValue string

		switch v := value.(type) {
		case int:
			encodedValue = fmt.Sprintf("%d", v)
		case uint64:
			encodedValue = fmt.Sprintf("%d", v)
		case int64:
			encodedValue = fmt.Sprintf("%d", v)
		case string:
			encodedValue = v
		case bool:
			encodedValue = fmt.Sprintf("%t", v)
		default:
			// handle other types if necessary
			continue
		}

		encodedParams = append(encodedParams, fmt.Sprintf("%s=%s", key, url.QueryEscape(encodedValue)))
	}

	return strings.Join(encodedParams, "&")
}

func (a *Ardor) GetRequest(path string) (ArdorResponse, error) {
	var ardorResp ArdorResponse

	timeout, _ := strconv.Atoi(httpTimeout)
	httpClient := http.Client{Timeout: time.Second * time.Duration(timeout)}

	req, err := http.NewRequest(http.MethodGet, a.buildURL(path), nil)
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
	if jsonErr != nil {
		return ardorResp, jsonErr
	}

	if ardorResp.ErrorCode != 0 {
		return ardorResp, fmt.Errorf("Ardor API returned an error: %s", ardorResp.ErrorDescription)
	}

	return ardorResp, nil
}

func (a *Ardor) PostRequest(data map[string]interface{}) (ArdorResponse, error) {
	var ardorResp ArdorResponse

	timeout, _ := strconv.Atoi(httpTimeout)
	httpClient := http.Client{Timeout: time.Second * time.Duration(timeout)}

	dd := encodeParams(data)
	req, err := http.NewRequest(http.MethodPost, a.Endpoint, strings.NewReader(dd))
	if err != nil {
		return ardorResp, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	if jsonErr != nil {
		return ardorResp, jsonErr
	}

	if ardorResp.ErrorCode != 0 {
		return ardorResp, fmt.Errorf("Ardor API returned an error: %s", ardorResp.ErrorDescription)
	}

	return ardorResp, nil
}

func (a *Ardor) GetRequestRaw(path string) ([]byte, error) {
	timeout, _ := strconv.Atoi(httpTimeout)
	httpClient := http.Client{Timeout: time.Second * time.Duration(timeout)}

	req, err := http.NewRequest(http.MethodGet, a.buildURL(path), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", appName, appVersion))

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(io.Reader(res.Body))
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}

func (a *Ardor) PostRequestRaw(data map[string]interface{}) ([]byte, error) {
	timeout, _ := strconv.Atoi(httpTimeout)
	httpClient := http.Client{Timeout: time.Second * time.Duration(timeout)}

	dd := encodeParams(data)
	req, err := http.NewRequest(http.MethodPost, a.Endpoint, strings.NewReader(dd))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", appName, appVersion))

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(io.Reader(res.Body))
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}
