package main

import (
	"net/http"

	"github.com/fransafu/chat-server/html"
)

func main() {
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		html.Chat(w, html.ChatParams{
			Title: "Chat",
		})
	})

	http.ListenAndServe(":8080", nil)
}
