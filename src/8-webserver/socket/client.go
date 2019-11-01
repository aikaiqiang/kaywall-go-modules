package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
	//	os.Exit(1)
	//}
	//service := os.Args[1]
	service := "127.0.0.1:1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	//checkError(err)
	// result, err := ioutil.ReadAll(conn)
	result := make([]byte, 256)
	_, err = conn.Read(result)
	checkError(err)
	fmt.Println(string(result))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: \n")
	// 单独开一个线程监听消息
	go printMsg(conn, result)

	for {
		text, _ := reader.ReadString('\n')
		if strings.Contains(text, "exit") {
			os.Exit(0)
		}
		_, err = conn.Write([]byte(text))
		checkError(err)
	}
}

func printMsg(conn *net.TCPConn, result []byte) {
	for {
		_, err := conn.Read(result)
		checkError(err)
		fmt.Println("Server:", string(result))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
