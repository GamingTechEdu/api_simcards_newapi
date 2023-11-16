package test_test

import (
	"bytes"
	"kdl_api_rest_internal/endpoint/post"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordSimcard(t *testing.T) {
	// create a request with a JSON body
	requestBody := []byte(`{
		"Client": "John Doe",
		"Iccid": "1234567890123456789",
		"Simcon": "2222002",
		"Msisdn": "1234567890",
		"Ip": "192.168.0.1",
		"Slot": "1",
		"Installationdate": "2022-01-01",
		"Activationdate": "2022-01-01",
		"Supplier": "Supplier",
		"Operator": "Operator",
		"Plan": "Plan",
		"Apn": "Apn",
		"Obs": "Obs"
	}`)

	req, err := http.NewRequest("POST", "/recordSimcard", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	// set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// call the handler function
	handler := http.HandlerFunc(post.RecordSimcard)
	handler.ServeHTTP(rr, req)

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the response body
	expected := "Dados gravados com sucesso!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
