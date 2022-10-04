package main

import (
	"log"
	"net/http"

	"github.com/fransafu/chat-server/html"
	"github.com/fransafu/chat-server/internal/chat_server"
)

func main() {
	hub := chat_server.NewHub()
	go hub.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html.Chat(w, html.ChatParams{
			Title: "Chat Server",
		})
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat_server.ServeWs(hub, w, r)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
