package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func parse(conn net.Conn) {
	var temp [128]byte
	for {
		n, err := conn.Read(temp[:])
		if err != nil {
			fmt.Println("read from conn failed err")
			return
		}
		fmt.Println(string(temp[:n]))
	}
}
func main() {
	//1. 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial failed")
		return
	}
	go parse(conn)
	//2. 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}
		conn.Write([]byte(text))
	}

	conn.Close()
}
