package main

import (
	"encoding/json"
	"net/http"
)

// Post struct
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{ID: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marchalling the posts array"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marchalling the request"}`))
		return
	}
	post.ID = len(posts) + 1
	posts = append(posts, post)
	response.WriteHeader(http.StatusOK)
	result, err := json.Marshal(posts)
	response.Write(result)

}
