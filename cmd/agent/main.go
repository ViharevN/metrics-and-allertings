package main

import (
	"metrics/pkg/httpclient"
	"time"
)

func main() {
	httpclient.NewAgent("http://127.0.0.1:8080", 2*time.Second, 10*time.Second).Run()
}
