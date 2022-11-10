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
	r.HandleFunc("/all/", controller.ReturnAllArticles)
	r.HandleFunc("/article/", controller.CreateNewArticle).Methods("POST")
	r.HandleFunc("/article/{id}", controller.UpdateArticle).Methods("PUT")
	r.HandleFunc("/article/{id}", controller.DeleteArticle).Methods("DELETE")
	r.HandleFunc("/article/{id}", controller.ReturnSingleArticle)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":2550", nil))
}
