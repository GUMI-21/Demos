package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	Address := "127.0.0.1:5000"
	conn, err := net.Dial("tcp", Address)
	if err != nil {
		fmt.Printf("dial %s failed,err:%s", Address, err)
		return
	}
	fmt.Println("客户端开启")
	var msg string
	var data [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入")
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		_, _ = conn.Write([]byte(msg))

		n, err := conn.Read(data[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("收到回复：", string(data[:n]))
	}
	_ = conn.Close()
}
