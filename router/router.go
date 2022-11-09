package router

import (
	"log"
	"net/http"

	"github.com/Isma1306/go-api/controller"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controller.HomePage)
	r.HandleFunc("/all", controller.ReturnAllArticles)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":2550", nil))
}
