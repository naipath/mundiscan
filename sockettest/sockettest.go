package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.171:1470")
	if err != nil {
		fmt.Print(err)
	}

	for {
		status, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Got the following data:", status)
	}
}
