package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"netCode/proto"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		log.Println("read from client failed, err:", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//var buf [1024]byte
	for {
		//n, err := reader.Read(buf[:])
		msg, err:= proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("read from client failed, err :", err)
			break
		}
		//recvStr := string(buf[:n])
		//log.Println("收到client发来的数据：", recvStr)
		log.Println("收到client发来的数据：", msg)
	}

}
