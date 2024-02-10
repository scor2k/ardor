package ardor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
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

func encodedStruct(data interface{}) string {
	var encodedParams []string

	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name
		fieldType := v.Type().Field(i).Type

		var encodedValue string

		switch fieldType.Kind() {
		case reflect.Int:
			encodedValue = fmt.Sprintf("%d", field.Int())
		case reflect.Uint64:
			encodedValue = fmt.Sprintf("%d", field.Int())
		case reflect.Int64:
			encodedValue = fmt.Sprintf("%d", field.Int())
		case reflect.String:
			encodedValue = field.String()
		case reflect.Bool:
			encodedValue = fmt.Sprintf("%t", field.Bool())
		default:
			// handle other types if necessary
			continue
		}

		encodedParams = append(encodedParams, fmt.Sprintf("%s=%s", fieldName, url.QueryEscape(encodedValue)))
	}

	return strings.Join(encodedParams, "&")
}

func (a *Ardor) GetRequestUnparsed(path string) (ArdorJsonResponse, error) {
	var response ArdorJsonResponse

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

	body, readErr := io.ReadAll(io.Reader(res.Body))
	if readErr != nil {
		return nil, readErr
	}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return response, nil
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
	return ardorResp, nil
}

func (a *Ardor) PostRawRequest(path string, payload string, timeout int) ([]byte, error) {
	httpClient := http.Client{Timeout: time.Second * time.Duration(timeout)}

	req, err := http.NewRequest(http.MethodPost, a.buildURL(path), strings.NewReader(payload))
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

func (a *Ardor) PostRawJsonRequest(path string, data ArdorJsonResponse, timeout int) ([]byte, error) {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res, err := client.Post(a.buildURL(path), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(io.Reader(res.Body))
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a *Ardor) GetRawRequest(path string, timeout int, sigbroToken string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, a.buildURL(path), nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", appName, appVersion))
	if len(sigbroToken) > 0 {
		req.Header.Set("X-Sigbro-Token", sigbroToken)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	response, err := io.ReadAll(io.Reader(res.Body))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func PostRequest(url string, data ArdorRequest) (ArdorResponse, error) {
	var ardorResp ArdorResponse

	timeout, _ := strconv.Atoi(httpTimeout)
	httpClient := http.Client{Timeout: time.Second * time.Duration(timeout)}

	dd := encodedStruct(data)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(dd))
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

	return ardorResp, nil
}

func postUrlencodedRequest(url string, data url.Values, timeout int) (ArdorJsonResponse, error) {
	var response ArdorJsonResponse
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	res, err := client.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(io.Reader(res.Body))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
