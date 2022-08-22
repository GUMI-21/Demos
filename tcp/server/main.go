package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func processConn(conn net.Conn) {
	var data [1024]byte
	var msg string
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(data[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read from conn failed,err:%s", err)
			return
		}
		fmt.Printf("Access Info: %s", string(data[:n]))
		fmt.Printf("回复信息:")
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		_, _ = conn.Write([]byte(msg))
	}
	defer conn.Close()
}

func main() {
	Address := "127.0.0.1:5000"
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("start tcp server %s failed ,err : %s ", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务端开启")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed,err: %s", err)
			return
		}
		go processConn(conn)
	}
}
