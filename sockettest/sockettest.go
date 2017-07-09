package main

import (
	"log"
	"net"

	"github.com/mdlayher/ethernet"
	"github.com/mdlayher/raw"
)

func main() {
	// Select the eth0 interface to use for Ethernet traffic.
	ifi, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatalf("failed to open interface: %v", err)
	}

	c, err := raw.ListenPacket(ifi, 0xcccc)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer c.Close()

	addr := &raw.Addr{HardwareAddr: ethernet.Broadcast}
	result, err := c.WriteTo([]byte{0x01, 0x02}, addr)

	if err != nil {
		log.Print(err)
	}
	log.Print(result)
}
