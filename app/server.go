package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	helpers "github.com/MehniLozo/meme-red/helpers"
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

	addr := net.TCPAddr{}
	addr.Port = 6379
	addr.IP = net.IPv4(0, 0, 0, 0)

	//l, err := net.Listen("tcp", "0.0.0.0:6379")
	l, err := net.ListenTCP("tcp", &addr)

	if err != nil {
		fmt.Println("Error listening on tcp", err)
		helpers.Die("Error while listening")
	}
	defer l.Close()
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
