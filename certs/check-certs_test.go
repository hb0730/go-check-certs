package certs

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	certs, err := Check("hb0730.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", certs)
}
