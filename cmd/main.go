package main

import (
	"github.com/matheus-osorio/go-email-validator/pkg/router"
)

func main() {
	router.Setup()
	router.Start()
}
