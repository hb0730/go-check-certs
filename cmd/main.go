package main

import (
	"flag"
	"github.com/hb0730/go-check-certs/http"
)

var (
	server string
)

func init() {
	flag.StringVar(&server, "server", ":80", "http license addr")
}
func main() {
	err := http.Request(server)
	if err != nil {
		panic(err)
	}
}
