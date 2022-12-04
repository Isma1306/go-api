package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	db "github.com/Isma1306/go-api/model"
	"github.com/Isma1306/go-api/types"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// var articles = []types.Article{
// 	{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 	{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
// }

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	articlesArray := db.GetAllArticles()
	json.NewEncoder(w).Encode(articlesArray)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: single article")
	vars := mux.Vars(r)
	id := vars["id"]
	response, err := db.GetArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went really wrong!")
	}

	json.NewEncoder(w).Encode(response)

}
func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Create article")
	reqBody, _ := io.ReadAll(r.Body)
	var article types.Article
	json.Unmarshal(reqBody, &article)
	id := db.CreateArticle(article)
	article.Id = id
	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Delete article")
	vars := mux.Vars(r)
	id := vars["id"]
	response, err := db.DeleteArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went really wrong!")
	}
	if response.DeletedCount == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("that article doesnt exist")

	} else {
		json.NewEncoder(w).Encode("Article deleted!")
	}

}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Update article")
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := io.ReadAll(r.Body)
	var article types.Article
	json.Unmarshal(reqBody, &article)
	response, err := db.UpdateArticle(article, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went really wrong!")
	}
	if response.MatchedCount == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("that article doesnt exist")

	} else {
		json.NewEncoder(w).Encode(response)
	}
}
