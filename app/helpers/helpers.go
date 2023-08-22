package helpers

import (
	"fmt"
	"net"
	"os"
)

func log(msg string) {
	fmt.Println(msg)

}

func Die(msg string) {
	log(msg)
	os.Exit(1)
}

func DoThing(bind net.Conn, n int) {
	if n < 0 {
		log("Ran into some error")
		return
	}
	bind.Write([]byte("Wanna do this thing so far"))

}
