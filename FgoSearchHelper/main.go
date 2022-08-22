package main

import (
	"example.com/m/FgoSearchHelper/data"
)

func main() {
	conn := data.GetConn("mysql")
	defer conn.Close()
}
