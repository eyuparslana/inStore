package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSet(t *testing.T) {
	requestParams := map[string]string{
		"key":   "test-key",
		"value": "test-value",
	}

	wantStatus := 201
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/set", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}
	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	t.Log("Test Success.")
}

func TestSetKeyError(t *testing.T) {
	requestParams := map[string]string{
		"value": "test-value",
	}

	wantStatus := 400
	wantError := "The 'key' is required."
	var response map[string]string
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/set", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}
	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("TestSet Error when ioutil.ReadAll: ", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal("TestSet Error when response unmarshall: ", err)
	}

	if response["error"] != wantError {
		t.Fatalf("Expected to get error %s but instead got %s", wantError, response["error"])
	}

	t.Log("Test Success.")
}

func TestSetValueError(t *testing.T) {
	requestParams := map[string]string{
		"key": "test-key",
	}

	wantStatus := 400
	wantError := "The 'value' is required."
	var response map[string]string
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/set", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}
	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("TestSet Error when ioutil.ReadAll: ", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal("TestSet Error when response unmarshall: ", err)
	}

	if response["error"] != wantError {
		t.Fatalf("Expected to get error %s but instead got %s", wantError, response["error"])
	}

	t.Log("Test Success.")
}

func TestSetMethodNotAllowed(t *testing.T) {
	requestParams := map[string]string{
		"key":   "test-key",
		"value": "test-value",
	}

	wantStatus := 405
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/set", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}
	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	t.Log("Test Success.")
}

func TestGet(t *testing.T) {
	requestParams := map[string]string{
		"key": "test-key",
	}
	want := "test-value"
	wantStatus := 200
	var response map[string]string
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/get", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}
	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("TestSet Error when ioutil.ReadAll: ", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal("TestSet Error when response unmarshall: ", err)
	}

	if response["result"] != want {
		t.Fatalf("Expected to get result %s but instead got %s", want, response["result"])
	}

	t.Log("Test Success.")
}

func TestGetNotFoundError(t *testing.T) {
	requestParams := map[string]string{
		"key": "test-keys",
	}
	wantError := "The key 'test-keys' could not be found."
	wantStatus := 404
	var response map[string]string
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/get", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}
	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("TestSet Error when ioutil.ReadAll: ", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal("TestSet Error when response unmarshall: ", err)
	}

	if response["error"] != wantError {
		t.Fatalf("Expected to get error %s but instead got %s", wantError, response["error"])
	}

	t.Log("Test Success.")
}

func TestHome(t *testing.T) {
	wantStatus := 200
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}
	t.Log("Test Success.")
}

func TestFlush(t *testing.T) {
	wantStatus := 204
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/flush", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}
	t.Log("Test Success.")
}
