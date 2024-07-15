package main

import (
	"github.com/ntc-goer/parser-exercise/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}
