package main

import (
	"log"

	"github.com/cryptotweet.io/cmd/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
