package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlers.helloHandler)
	http.HandleFunc("/article", handlers.postArticleHandler)
	http.HandleFunc("/article/list", handlers.articleListHandler)
	http.HandleFunc("/article/1", handlers.articleDetailHandler)
	http.HandleFunc("/article/nice", handlers.postNiceHandler)
	http.HandleFunc("/comment",handlers.postCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}