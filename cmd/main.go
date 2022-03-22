package main

import (
	"log"

	"github.com/levinhne/cryptotweet.io/cmd/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
