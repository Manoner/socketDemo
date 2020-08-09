package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	log.Println("服务端启动....")
	// 监听
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	// 循环接收多个用户
	for {
		// 阻塞等待用户连接
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			return
		}

		// 获取客户端的网络地址信息
		addr := conn.RemoteAddr().String() // .String() 转换成字符串
		log.Printf("客户端 %s 连接成功 \n", addr)

		// 处理用户请求，每来一个用户就新建一个 go 协程
		go process(conn) // 启动一个goroutine处理连接
		log.Println("main for loop in ")
	}
	fmt.Printf("main for loop out")
}

// 处理用户请求
func process(conn net.Conn) { // conn 是 net.Conn 类型
	// 函数调用完毕，自动关闭 conn
	defer conn.Close()
	// 获取客户端的网络地址信息
	addr := conn.RemoteAddr().String() // .String() 转换成字符串
	log.Println(addr, " connect successful")

	for {
		// 接收用户的请求
		reader := bufio.NewReader(conn)
		var buf [128]byte             // 声明切片缓冲区，其大小为128
		n, err := reader.Read(buf[:]) // 读取数据放入buf，返回读取数据的长度
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}

		// 读了多少字节数据就打印多少数据
		recvStr := string(buf[:n])
		fmt.Printf("收到client端 %s 发来的数据：%s", addr, recvStr)

		if "Q" == string(buf[:n-1]) {
			log.Println(addr, "  exit ")
			return
		}

		// 向客户端返回读取到的数据
		recvStrUpper := strings.ToUpper(recvStr) // 将字符转换为大写
		conn.Write([]byte(recvStrUpper))
	}
	fmt.Println("服务器退出for循环")

}
