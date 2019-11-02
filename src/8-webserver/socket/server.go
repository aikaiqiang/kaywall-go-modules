package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//创建字典集合存储客户端连接
var conns = make(map[string]net.Conn)

func main() {
	//server_1()
	//server_2()
	server_3()

}

func server_1() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("port 7777 request coming ...")
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
	}
}

func server_2() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("port 1200 request coming ...")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime)) // don't care about return value
	// we're finished with this client
}

func server_3() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	log.Println("等待客户端连接")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		connIp := conn.RemoteAddr().String()
		// 保存连接
		conns[connIp] = conn
		log.Println(connIp, "已经建立连接")
		welcome := "welcome " + conn.RemoteAddr().String() + "\n"
		// server response
		conn.Write([]byte(welcome))
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			log.Println(conn.RemoteAddr().String(), "出现错误：", err)
			delete(conns, conn.RemoteAddr().String())
			broadMsg([]byte(conn.RemoteAddr().String() + "已下线"))
			return
		}

		msg := conn.RemoteAddr().String() + ":" + string(buffer[:n])
		broadMsg([]byte(msg))
		log.Println("用户【", conn.RemoteAddr().String(), "】发送一条消息", string(buffer[:n]))

	}
}

func broadMsg(msg []byte) { //消息广播
	for _, connopj := range conns {
		//log.Println("用户【", connopj.RemoteAddr().String(),"】发送一条消息",string(msg))
		connopj.Write(msg)
	}
}

func handleClientNew(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128B to prevent flood attack
	defer conn.Close()                                    // close connection before exit
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal 11-error: %s", err.Error())
		os.Exit(1)
	}
}
