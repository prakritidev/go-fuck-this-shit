// package main

// import (
//     "fmt"
//     "log"
//     // "os"
//     // "os/signal"
//     // "syscall"
//     // "time"

//     // "github.com/google/gopacket"
//     // "github.com/google/gopacket/layers"
//     // "github.com/google/gopacket/pcap"
//     // "github.com/google/gopacket/pfring"
// )

// func main() {
//     // Open a new packet capture handle on the specified interface
//     handle, err := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer handle.Close()

//     // Set filter to capture only TCP traffic
//     err = handle.SetBPFFilter("tcp")
//     if err != nil {
//         log.Fatal(err)
//     }

//     // Create packet source
//     packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

//     // Capture packets indefinitely
//     for packet := range packetSource.Packets() {
//         // Extract TCP layer
//         tcpLayer := packet.Layer(layers.LayerTypeTCP)
//         if tcpLayer != nil {
//             tcp, _ := tcpLayer.(*layers.TCP)

//             // Print source and destination port
//             fmt.Printf("Source Port: %d, Destination Port: %d\n", tcp.SrcPort, tcp.DstPort)
//         }
//     }
// }

