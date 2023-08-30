package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", webSocketHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected")

	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(c *websocket.Conn) {
	for {
		messageType, p, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		n, _ := strconv.Atoi(string(p))

		if err := c.WriteMessage(messageType, []byte(strconv.Itoa(n*n))); err != nil {
			log.Println(err)
			return
		}
	}
}
