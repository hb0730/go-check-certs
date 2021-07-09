package http

import "testing"

func TestRequest(t *testing.T) {
	err := Request(":80")
	if err != nil {
		panic(err)
	}
}
