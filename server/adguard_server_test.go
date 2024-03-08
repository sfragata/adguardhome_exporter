package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

const validJSON = "{\"MediaContainer\":{\"size\":1}}"

type response struct {
	MediaContainer struct {
		Size int `json:"size"`
	} `json:"MediaContainer"`
}

func newAdguardServer(server *httptest.Server) AdguardServer {

	hostPort := strings.Split(server.Listener.Addr().String(), ":")

	port, _ := strconv.Atoi(hostPort[1])

	return AdguardServer{
		Host: hostPort[0],
		Port: port,
	}

}
func TestGivenValidJsonResponseWhenSendResquestThenReturnValidStruct(test *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validJSON)
	}))
	defer server.Close()

	AdguardServer := newAdguardServer(server)

	var response response
	err := AdguardServer.SendRequest("api", &response)
	if err != nil {
		test.Errorf("Error: %v", err)
	}
	if response.MediaContainer.Size != 1 {
		test.Errorf("Limits: \nExpected: %d\nActual: %d", 1, response.MediaContainer.Size)
	}
}

func TestGivenAdguardTokenWhenSendResquestThenAdguardTokenIsInHeader(test *testing.T) {

	expectedToken := "1234"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adguardToken := r.Header.Get("Authorization")
		if adguardToken != fmt.Sprintf("Basic %s", expectedToken) {
			test.Errorf("\nExpected: %s\nActual: %s", expectedToken, adguardToken)
		}
		fmt.Fprint(w, validJSON)

	}))
	defer server.Close()

	AdguardServer := newAdguardServer(server)
	AdguardServer.Token = expectedToken

	var response response
	err := AdguardServer.SendRequest("api", &response)
	if err != nil {
		test.Errorf("Error: %v", err)
	}
}

func TestGivenServerErrorWhenSendResquestThenReturnError(test *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer server.Close()

	AdguardServer := newAdguardServer(server)
	var response response
	err := AdguardServer.SendRequest("api", &response)

	if err == nil {
		test.Error("Should throw error")
	}
	errorMsg := fmt.Sprintf("%v", err)
	if !strings.Contains(errorMsg, "error: status code 500 from server") {
		test.Errorf("Error, it should contains expected: 'Error: status code 500 from server', actual: '%v'", err)
	}
}

func TestGivenInvalidJsonResponseWhenSendRequestThenReturnError(test *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "invalidJSON")
	}))
	defer server.Close()

	AdguardServer := newAdguardServer(server)
	var response response
	err := AdguardServer.SendRequest("api", &response)

	if err == nil {
		test.Errorf("Error: %v", err)
	}
	errorToString := fmt.Sprintf("%v", err)
	if !strings.Contains(errorToString, "invalid JSON") {
		test.Errorf("\nShould contain: %s\nActual: %s", "Invalid JSON", errorToString)
	}

}
