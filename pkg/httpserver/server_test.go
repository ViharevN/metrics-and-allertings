package httpserver

import (
	"net/http"
	"testing"
)

func TestServer_ErrServ(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := New(handler)

	client := &http.Client{}

	res, err := client.Post("http://"+server.server.Addr, "text/plain", nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", res.StatusCode)
	}

}
