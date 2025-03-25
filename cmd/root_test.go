package cmd

import (
	"github.com/lokalise/go-lokalise-api/v4"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	testApiToken  = "apiToken"
	testProjectID = "3002780358964f9bab5a92.87762498"
)

func setup() (api *lokalise.Api, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()

	server := httptest.NewServer(mux)

	apiURL, _ := url.Parse(server.URL + "/api2/")
	api, _ = lokalise.New(testApiToken, lokalise.WithBaseURL(apiURL.String()))

	return api, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is \n%s\n want\n%s", got, want)
	}
}
