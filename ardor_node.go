package ardor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type nodeStatus struct {
	up             bool
	time           uint64
	numberOfBlocks uint64
	version        string
	responseTimeMs int64
}

func getNodeStats(_host string, _type string) (stat nodeStatus) {
	url := fmt.Sprintf("http://%s/nxt?requestType=getBlockchainStatus", _host)

	client := http.Client{
		Timeout: 1 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		return nodeStatus{up: false}
	}

	body, err := io.ReadAll(io.Reader(res.Body))
	if err != nil {
		return nodeStatus{up: false}
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nodeStatus{up: false}
	}
	version := fmt.Sprintf("%s", response["version"])
	version = strings.TrimSpace(version)

	stat.version = version

	_numOfBlocks := fmt.Sprintf("%.0f", response["numberOfBlocks"])
	stat.numberOfBlocks, _ = strconv.ParseUint(_numOfBlocks, 10, 64)

	_time := fmt.Sprintf("%.0f", response["time"])
	stat.time, _ = strconv.ParseUint(_time, 10, 64)

	blockchainState := response["blockchainState"]
	if blockchainState == "UP_TO_DATE" {
		stat.up = true
		return
	}
	if blockchainState == "FORK" {
		stat.up = true
		return
	}

	stat.up = false
	return
}

func (a *Ardor) getBlockchainStatus(network string) (ArdorBlockchainStatusResponse, error) {
	payload := "requestType=getBlockchainStatus"

	response, err := a.PostRawRequest(a.Endpoint, payload, 10)
	if err != nil {
		return ArdorBlockchainStatusResponse{}, err
	}

	var jsonResponse ArdorBlockchainStatusResponse
	// parse response
	jsonErr := json.Unmarshal(response, &jsonResponse)
	if jsonErr != nil {
		return ArdorBlockchainStatusResponse{}, jsonErr
	}

	if jsonResponse.NumberOfBlocks < 1 {
		msg := fmt.Sprintf("getBlockchainStatus status is wrong: %s, height: %v", jsonResponse.BlockchainState, jsonResponse.NumberOfBlocks)
		return ArdorBlockchainStatusResponse{}, fmt.Errorf("%s", msg)
	}

	return jsonResponse, nil
}
