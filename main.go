package main

import (
	"fmt"
	"go-eth-auction/http"
)

const (
	Port = 8087
)

func main() {
	var s, err = http.NewServer(Port)

	if err != nil {
		panic(err)
	}

	fmt.Printf("trying to start http server on port %d\n", Port)
	if err := s.Start(); err != nil {
		fmt.Printf("failed to run http server, error: %s\n", err.Error())
	}
}
