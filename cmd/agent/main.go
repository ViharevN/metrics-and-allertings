package main

import (
	"metrics/pkg/httpclient"
	"time"
)

func main() {
	httpclient.NewAgent("http://localhost:8080", 2*time.Second, 10*time.Second).Run()
}
