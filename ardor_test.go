package ardor

import "testing"

func TestArdor_GetRequest(t *testing.T) {
	ardor := Ardor{Endpoint: "https://random.api.nxter.org/ardor"}

	data, err := ardor.GetRequest("?requestType=getTime")
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
		return
	}

	if data.ErrorCode != 0 {
		t.Errorf("Expected ErrorCode to be 0, but got %d", data.ErrorCode)
	}

	if data.Time == 0 {
		t.Errorf("Expected Time to be a non-zero value, but got 0")
	}
}

func TestArdor_PostRequest(t *testing.T) {
	ardor := Ardor{Endpoint: "https://random.api.nxter.org/ardor"}

	params := map[string]interface{}{
		"requestType": "getTime",
	}
	data, err := ardor.PostRequest(params)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
		return
	}

	if data.ErrorCode != 0 {
		t.Errorf("Expected ErrorCode to be 0, but got %d", data.ErrorCode)
	}

	if data.Time == 0 {
		t.Errorf("Expected Time to be a non-zero value, but got 0")
	}
}
