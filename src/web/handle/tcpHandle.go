package handle

import (
	"log"
	"net"
	"wechat/src/web/chat/discussMsg"
)

//var param = make(chan []byte,3)
var buffer = make([]byte, 1024)
var link = make(map[string]net.Conn)

func Tcp(listener net.Listener) {

	for {
		conn, err := listener.Accept()
		link[conn.RemoteAddr().String()] = conn
		if err != nil {
			log.Println("error", err)
			continue
		}
		n, err := conn.Read(buffer)

		result := getResult(discussMsg.AddDiscussMsg(buffer[:n]))
		go writeResult(result)

	}

}
func writeResult(result []byte) {
	for _, conns := range link {
		_, _ = conns.Write(result)
	}
}
