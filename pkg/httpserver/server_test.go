package httpserver

import (
	"net/http"
	"testing"
	"time"
)

func TestServer_ErrServ(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := New(handler)

	time.Sleep(100 * time.Millisecond)

	client := &http.Client{}

	res, err := client.Post("http://"+server.server.Addr, "text/plain", nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", res.StatusCode)
	}

}
