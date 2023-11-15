package main

import (
	"fmt"
	wstrying "wsTrying"
)

func main() {
	fmt.Println("Hello")
	server := wstrying.StartServer(HandleMessage)
	var command string
	for {
		fmt.Scanln(&command)
		if command == "shutdown" {
			break
		}
		switch command {
		case "lookup":
			{
				fmt.Println(server)
			}
		}
	}
}

func HandleMessage(message wstrying.Message) {
	fmt.Println(string(message.Msg))
	message.Session.SendMessage(message.Mt, message.Msg)
}
