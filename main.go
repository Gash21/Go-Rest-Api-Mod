package main

import (
	"ModRestApi/app/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/article/{id}", singleArticle)
	myRouter.HandleFunc("/create", createNewArticle)
	myRouter.HandleFunc("/update", updateArticle)
	log.Fatal(http.ListenAndServe(":8050", myRouter))
}

func singleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint single article")
	vars := mux.Vars(r)
	key := vars["id"]

	json.NewEncoder(w).Encode(model.FindArticle(key, model.Articles))
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint all articles")
	data := map[string]interface{}{
		"articles": model.Articles,
	}

	result := map[string]interface{}{
		"message": "success",
		"status":  200,
		"data":    data,
	}

	b, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	model.CreateNewArticle(
		model.Article{
			Id:      r.Form["Id"][0],
			Title:   r.Form["Title"][0],
			Desc:    r.Form["Desc"][0],
			Content: r.Form["Content"][0],
		},
	)

	msg := model.Message{Message: "success", Status: "200"}
	b, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	artId := r.Form["Id"][0]
	msg := model.Message{Message: "success", Status: "200"}

	if artId == "0" {
		msg = model.Message{Message: "not found", Status: "400"}
	} else {
		model.UpdateArticle(
			model.Article{
				Id:      r.Form["Id"][0],
				Title:   r.Form["Title"][0],
				Desc:    r.Form["Desc"][0],
				Content: r.Form["Content"][0],
			},
		)
	}

	b, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	model.Populate()
	handleRequests()
}
