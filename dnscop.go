package main

import (
	"dnscop/block"
	"dnscop/config"
	"dnscop/dnsmsg"
	"flag"
	"log"
	"net"
	"strings"
)

const maxDnsPacketSize = 512
const VERSION = "0.0.1"

func main() {
	listen := flag.String("listen", ":53", "listen IP:Port. ex. 127.0.0.1:53")
	resolver := flag.String("resolver", "8.8.8.8:53", "Resolver IP:Port. ex. 8.8.8.8:53")
	between := flag.String("between", "20:00-23:00", "between time ex. 20:00-07:00")
	blockList := flag.String("block-list", "www.youtube.com|youtube.com|i.ytimg.com|.+.googlevideo.com", "block list domains. ")
	flag.Parse()

	log.Printf("==== DNSCOP (ver %s) ====\n", VERSION)
	log.Println("Listen: " + *listen)
	log.Println("Resolver: " + *resolver)
	log.Println("Between: " + *between)
	log.Println("BlockList: " + *blockList)

	packet, err := net.ListenPacket("udp", *listen)
	if err != nil {
		log.Fatal(err)
	}
	defer packet.Close()

	conf := config.NewUserConfig(*blockList, *between)

	for {
		buf := make([]byte, maxDnsPacketSize)
		readbyte, clientAddr, err := packet.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		go handleDnsRequest(packet, clientAddr, buf[:readbyte], *resolver, conf)
	}
}

func handleDnsRequest(packet net.PacketConn, address net.Addr, data []byte, resolver string, conf *config.UserConfig) {
	//log.Println(data)
	name, err := dnsmsg.GetQuestionName(data)
	name = strings.TrimRight(name, ".")

	log.Println(name + " " + address.String())

	if block.IsBlock(name, conf) {
		log.Println("  ** block: " + name + " **  " + address.String())
		return
	}

	response, err := dnsmsg.Send(resolver, data)
	if err != nil {
		log.Fatal(err)
	}
	packet.WriteTo(response, address)
}
