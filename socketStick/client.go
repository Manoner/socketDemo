package main

import (
	"log"
	"net"
	"netCode/proto"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		log.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := `Hello,Hello. How are you?`
		//conn.Write([]byte(msg))

		data, err := proto.Encode(msg)
		if err != nil {
			log.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
