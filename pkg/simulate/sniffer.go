package simulate

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"k8s.io/klog/v2"
)

func chooseNetworkDevice() (string, error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return "", fmt.Errorf("find device: %w", err)
	}

	foundAddrs := 0
	foundName := ""

	for _, d := range devices {
		if strings.HasPrefix(d.Name, "en") || strings.HasPrefix(d.Name, "wl") {
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
	klog.Infof("looking for network devices ...")
	iface, err := chooseNetworkDevice()
	if err != nil {
		return err
	}

	klog.Infof("chose %s", iface)
	klog.Infof("opening pcap ...")

	handler, err := pcap.OpenLive(iface, 1600, false, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("openlive: %w", err)
	}
	defer handler.Close()

	klog.Infof("setting filter ...")
	if err := handler.SetBPFFilter("ip"); err != nil {
		log.Fatal(err)
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	klog.Infof("iterating over packets ...")
	for p := range source.Packets() {
		klog.Infof("got packet: %+v", p)
		return nil
	}
	return nil
}
