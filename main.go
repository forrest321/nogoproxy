package main

import (
	"gopkg.in/elazarl/goproxy.v1"
	"log"
	"net/http"
)

func main() {
	//1. parse hosts file to fastest datastore availavble
	//2. accept any request not in hosts file
	//4. profit
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
