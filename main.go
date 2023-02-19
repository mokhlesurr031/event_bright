package main

import (
	"github.com/event_bright/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
	}
}
