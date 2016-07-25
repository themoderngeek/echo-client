package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//service := "localhost:2020"
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	conn, err := net.Dial("tcp", "127.0.0.1:2020")
	if err != nil {
		println(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		println(err)
	}
	fmt.Fprintf(conn, text)

	buf := make([]byte, 1024)
	inputLength, err := conn.Read(buf)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Message recived %s", buf[0:inputLength])
}
