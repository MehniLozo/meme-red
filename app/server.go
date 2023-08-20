package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func requestReponse(bind net.Conn) {
	name := bind.RemoteAddr().String()
	fmt.Println("Receiving request from ", name)
	message := bufio.NewScanner(bind)
	message.Scan()
	fmt.Printf("%s: %s\n", name, message.Text())
	bind.Write([]byte("+PONG\r\n"))
}
func main() {
	fmt.Println("Logs from a very professional redis")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	closed := false
	for !closed {
		bind, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			bind.Close()
			closed = true
			os.Exit(1)
		}
		go requestReponse(bind)
	}
}
