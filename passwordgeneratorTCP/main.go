package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

var char = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"@", "#", "$", "&"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	portptr := flag.Int("port", 8080, "listening port")
	flag.Parse()
	port := strconv.Itoa(*portptr)

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go generatePassword(conn)
	}
}

func generatePassword(conn net.Conn) {
	log.Println("Start password generator")
	passwd := ""
	for i := 0; i < 8; i++ {
		passwd = passwd + char[rand.Intn(65)]
	}
	fmt.Fprintf(conn, "%s\n", passwd)
	log.Printf("Generated password = %s\n", passwd)
	log.Println("End password generator")
}
