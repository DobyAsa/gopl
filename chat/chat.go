package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Chat已监听在8000端口")

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

type client struct {
	chat chan<- string
	addr string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.chat <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			log.Printf("%s 加入聊天", cli.name)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.chat)
			log.Printf("%s 离开聊天", cli.name)
		}

	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 20)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	name := ""
	input := bufio.NewScanner(conn)
	io.WriteString(conn, "请输入你的昵称:")
	if input.Scan() {
		name = input.Text()
	}
	cli := client{
		chat: ch,
		addr: who,
		name: name,
	}
	ch <- "You are " + cli.name
	messages <- cli.name + " 进入聊天室"
	entering <- cli

	for input.Scan() {
		if cli.name == "" {
			messages <- who + ": " + input.Text()
		}
		messages <- cli.name + ": " + input.Text()
	}

	leaving <- cli
	messages <- cli.name + " 离开聊天室"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, "\r%s\n", msg)
	}
}
