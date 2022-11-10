package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

var articles = []Article{
	{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: single article")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}
func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Create article")
	reqBody, _ := io.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Delete article")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
			json.NewEncoder(w).Encode(article)
		}
	}

}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Update article")
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := io.ReadAll(r.Body)
	var new Article
	json.Unmarshal(reqBody, &new)
	for index, article := range articles {
		if article.Id == id {
			articles[index] = new
			json.NewEncoder(w).Encode(article)
		}
	}

}
