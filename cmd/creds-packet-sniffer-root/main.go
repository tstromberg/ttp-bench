// Simulates theft of credentials via network sniffing [T1040]
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	if err := PacketSniffer(); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}

func chooseNetworkDevice() (string, error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return "", fmt.Errorf("find device: %w", err)
	}

	foundAddrs := 0
	foundName := ""

	for _, d := range devices {
		if strings.HasPrefix(d.Name, "en") || strings.HasPrefix(d.Name, "eth") || strings.HasPrefix(d.Name, "wl") {
			if len(d.Addresses) > foundAddrs {
				foundName = d.Name
				foundAddrs = len(d.Addresses)
			}
		}
	}
	if foundAddrs > 0 {
		return foundName, nil
	}

	return "", fmt.Errorf("could not pick a device among: %v", devices)
}

func PacketSniffer() error {
	log.Printf("looking for network devices ...")
	iface, err := chooseNetworkDevice()
	if err != nil {
		return err
	}

	log.Printf("chose %s", iface)
	log.Printf("opening pcap ...")

	handler, err := pcap.OpenLive(iface, 1600, false, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("openlive: %w", err)
	}
	defer handler.Close()

	log.Printf("setting filter ...")
	if err := handler.SetBPFFilter("ip"); err != nil {
		log.Fatal(err)
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	log.Printf("iterating over packets ...")
	packets := 0
	for p := range source.Packets() {
		log.Printf("got packet: %+v", p)
		packets++
		if packets >= 2 {
			break
		}
	}
	return nil
}
