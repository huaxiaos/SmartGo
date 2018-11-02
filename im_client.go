package main

import (
	"bufio"
	"fmt"
	"net"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)

	//b := []byte("time\n")
	//conn.Write(b)

	// 控制台聊天功能加入
	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "quit" {
			break
		}
		b := []byte(msg + "\n")
		conn.Write(b)
	}

	<-quitSemaphore
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		//time.Sleep(time.Second)
		//b := []byte(msg)
		//conn.Write(b)
	}
}