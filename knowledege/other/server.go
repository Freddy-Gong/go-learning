package knowledege

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func process(id int, conn net.Conn) {
	var temp [128]byte
	for {
		n, err := conn.Read(temp[:])
		if err != nil {
			fmt.Println("read from conn failed err")
			return
		}
		//给其他连接发消息
		for i, con := range clientList {
			if i == id {
				continue
			}
			con.Write(temp[:n])
		}
		fmt.Println(string(temp[:n]))
	}

}

var clientList map[int]net.Conn

func Server() {
	clientList = make(map[int]net.Conn, 100)
	//1. 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server failed")
		return
	}
	id := 0
	for { //客户端和服务端需要连接多次 所以要用for循环
		//2. 等待建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed")
			return
		}
		clientList[id] = conn
		//3. 客户端通信 可开启其他线程进行处理
		go process(id, conn)
		id++
	}
}

//如何解决粘包 规定前四个字节为包的长度
func Encode(msg string) ([]byte, error) {
	//读取消息的长度 转换成int32类型 占4个字节
	length := int32(len(msg))
	pkg := new(bytes.Buffer) //新建一个缓冲区
	//写入消息头 binary以字节的形式进行操作
	//binary.LittleEnddian 小端写入 倒着写
	//binary.BigEnddian 打端写入 正着写
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	//读取消息的长度 读取前四个字节
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Write(lengthBuff, binary.LittleEndian, length)
	if err != nil {
		return "", err
	}
	//buffered返回缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	//读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

//udp server端
func UdpServer() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen udp faile")
		return
	}
	//udp就不需要建立连接了 直接收发数据
	var data [1024]byte
	for {
		//接收消息的时候就可以得到发送端的地址了
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from udp err")
			return
		}
		fmt.Println(data[:n])
		//发送数据
		replay := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(replay), addr)
	}
}
