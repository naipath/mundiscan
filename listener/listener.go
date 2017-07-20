package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.171:1471")

	if err != nil {
		fmt.Print("error opening connection\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	for {
		result, _ := bufio.NewReader(conn).ReadBytes(0x04)

		fmt.Printf("Received:\n%08b\n", result)

		conn.Write(getStatusMessageResponse())
	}
}

func getVersionMessageResponse() []byte {
	lsb, msb := getCheckSum(0x56, 0x06, 0x32, 0x31, 0x30, 0x30, 0x30, 0x35)
	return []byte{0x02, 0x56, 0x06, 0x32, 0x31, 0x30, 0x30, 0x30, 0x35, lsb, msb, 0x04}
}

func getCurrentCountResponse() []byte {
	lsb, msb := getCheckSum(0x48, 0x08, 0, 0, 0, 42, 0, 0, 0, 23)
	return []byte{0x02, 0x48, 0x08, 0, 0, 0, 42, 0, 0, 0, 23, lsb, msb, 0x04}
}

func getStatusData() []byte {
	lsb, msb := getCheckSum(0x5A, 0x06, 192, 255, 43, 3, 224)
	return []byte{0x02, 0x5A, 0x06, 0, 1, 192, 255, 43, 3, 224, lsb, msb, 0x04}
}

func getStatusMessageResponse() []byte {
	lsb, msb := getCheckSum(0x5B, 0x06, 0x00, 0x41, 0x00, 0x42, 0x00, 0x43)
	return []byte{0x02, 0x5B, 0x06, 0x00, 0x41, 0x00, 0x42, 0x00, 0x43, lsb, msb, 0x04}
}

func getCheckSum(input ...byte) (byte, byte) {
	var total uint16
	for _, value := range input {
		total += uint16(value)
	}
	lsb := byte(total & 0xFF)
	msb := byte((total >> 8) & 0xFF)
	return lsb, msb
}
