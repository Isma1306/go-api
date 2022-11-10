package main

import (
	"fmt"

	"github.com/Isma1306/go-api/router"
)

func main() {
	fmt.Println("Api started!")
	router.HandleRequests()
}
