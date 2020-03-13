package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/luisabarbalho/go_cache_pragmatic_review/entity"
	"github.com/luisabarbalho/go_cache_pragmatic_review/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRespository()
)

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := repo.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marchalling the request"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
