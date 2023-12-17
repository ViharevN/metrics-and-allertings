package main

import (
	"metrics/pkg/consts"
	"metrics/pkg/httpclient"
	"time"
)

func main() {
	httpclient.NewAgent("http://"+consts.AddrClient, 2*time.Second, 10*time.Second).Run()
}
