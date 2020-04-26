package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article // array of struct Article

var counter int

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit : updateArticle")
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, id)
	var articleUpdate Article
	json.Unmarshal(reqBody, &articleUpdate)

	for index, article := range Articles {
		if article.Id == id {
			Articles[index].Content = articleUpdate.Content
			Articles[index].Desc = articleUpdate.Desc
			Articles[index].Title = articleUpdate.Title
			fmt.Fprintf(w, "Updated successfully")
			return
		}
	}
	fmt.Fprintf(w, "Invalid id")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit : deleteArticle")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			return
		}
	}
	fmt.Fprintf(w, "Invalid id")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit : createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
	// fmt.Fprintf(w, "%+v", string(reqBody))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit : returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit : returnSingleArticle")

	vars := mux.Vars(r)
	key := vars["id"]

	// fmt.Fprintf(w, "Key: "+key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	fmt.Fprintf(w, "Invalid id")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
	fmt.Println("End point hit: homepage", counter)
	counter++

}

func handleRequests() {
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Rest API V2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
