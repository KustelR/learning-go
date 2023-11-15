package wstrying

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	sessions Sessions
	//handleMessage func(message Message) // хандлер новых сообщений
}

func StartServer(handleMessage func(message Message)) *Server {
	server := Server{
		Create(),
	}

	http.HandleFunc("/", server.handleReq)
	go http.ListenAndServe(":8080", nil) // Уводим http сервер в горутину

	return &server
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

func (server *Server) handleReq(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	go server.init(connection)
	//fmt.Println(r)
	//id, _ := server.sessions.CreateSession([]*websocket.Conn{connection}) // Сохраняем соединение, используя его как ключ
	//go server.sessions.AddMessageHandler(id, server.handleMessage)
	//connection.WriteMessage(websocket.TextMessage, []byte(id))
}

func (server *Server) init(connection *websocket.Conn) {
	fmt.Println("Detected connection")
	for {
		_, msg, _ := connection.ReadMessage()
		message := string(msg)
		if string(msg) == "init" {
			fmt.Println("started initialization")
			id, _ := server.sessions.CreateSession([]*websocket.Conn{connection})
			connection.WriteMessage(websocket.TextMessage, []byte(id))
			fmt.Printf("initialized: %v session\n", id)
			go server.sessions.AddMessageHandler(id, echo)
			break
		} else if server.sessions.FindSession(message) != nil {
			server.sessions.AddListener(message, connection)
			fmt.Printf("connected to %v session\n", message)
			break
		}

	}
}

func echo(message Message) {
	fmt.Println(message)
	fmt.Printf("echoing from %v to %v\n", message.Session, message.Session.listeners)
	message.Session.SendMessage(message.Mt, message.Msg)
}
