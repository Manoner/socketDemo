package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	log.Println("客户端程序启动...")
	// 主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("client Dial failed, err:", err)
		return
	}
	defer conn.Close() // 关闭连接

	// 从键盘读取用户输入，并存储在 inputReader 变量
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // ，回车结束
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果用户输入Q，就退出
			log.Println("退出成功")
			return
		}
		// 把输入的内容给服务器发送
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			log.Println("write err ,err :", err)
			return
		}

		// 从服务器读取
		buf := [512]byte{} // 切片缓冲
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed,err:", err)
			return
		}
		log.Println("recv from server: ",string(buf[:n])) // 打印接收到的内容，转换为字符再打印
	}

}
