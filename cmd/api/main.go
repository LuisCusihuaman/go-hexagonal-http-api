package main

import (
	"github.com/LuisCusihuaman/go-hexagonal-http-api/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
