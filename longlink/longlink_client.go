package longlink

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var quitSemaphoreLongLink chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecivedLongLink(conn)

	b := []byte("time\n")
	conn.Write(b)

	<-quitSemaphoreLongLink
}

func onMessageRecivedLongLink(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphoreLongLink <- true
			break
		}
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write(b)
	}
}