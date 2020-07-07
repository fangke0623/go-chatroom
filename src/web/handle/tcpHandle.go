package handle

import (
	"log"
	"net"
	"wechat/src/web/chat/discussMsg"
)

//var param = make(chan []byte,3)
var buffer = make([]byte, 1024)
var linkGroup = make(map[string]net.Conn)

func Tcp(listener net.Listener) {

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error", err)
			continue
		}
		go handleConn(conn)
	}

}
func handleConn(conn net.Conn) {
	linkGroup[conn.RemoteAddr().String()] = conn

	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("handleConn", err)
	}
	result := getResult(discussMsg.AddDiscussMsg(buffer[:n]))
	writeResult(result)
}

func writeResult(result []byte) {
	for _, conn := range linkGroup {
		go conn.Write(result)
	}
}
