// +build !solution

// Leave an empty line above this comment.
//
// Zap Collection Server
package main

import (
	"encoding/hex"
	"log"
	"net"
	"os"

	"github.com/uis-dat320-fall16/glabs/lab7/zlog"
)

const (
	clientAddr      = "224.0.0.1:1000"
	maxDatagramSize = 8192
)

// REMARK: This function should return (i.e. it should not block)
func runLab() {
	switch *labnum {
	case "a", "c1", "c2", "d", "e":
		ztore = zlog.NewSimpleZapLogger()
	case "f":
		//TODO activate with new ZapLogger data structure (task f)
		// ztore = zlog.NewViewersZapLogger()
	}
	switch *labnum {
	case "a":
		//TODO write code for dumping zap events to console
		// go dumpAll()
	case "c1":
		//TODO write code for recording and showing # of viewers on NRK1
		// go recordAll()
		// go showViewers("NRK1")
	case "c2":
		//TODO write code for task c2
	case "d":
		//TODO write code for task d
	case "e":
		//TODO write code for task e
	case "f":
		//TODO write code for task f
	}
}

// REMARK: This function should return (i.e. it should not block)
func startServer() {
	log.Println("Starting ZapServer...")
	//TODO write this method (5p)

	udpAddr, err := net.ResolveUDPAddr("udp", clientAddr)
	checkForErr(err)
	l, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	checkForErr(err)

	l.SetReadBuffer(maxDatagramSize)
	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		msgHandler(src, n, b)
	}

}

func checkForErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}
