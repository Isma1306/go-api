package main

import (
	"fmt"

	db "github.com/Isma1306/go-api/model"
	"github.com/Isma1306/go-api/router"
	"github.com/Isma1306/go-api/utils"
)

var Username string
var Password string

func main() {
	utils.Prompt()
	db.Connect()
	fmt.Println("Api started!")
	router.HandleRequests()
}
