package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	portptr := flag.Int("port", 8080, "Port of password generator server")
	hostptr := flag.String("ip", "127.0.0.1", "IP address of password generator server")
	flag.Parse()

	match, err := regexp.MatchString("^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$", *hostptr)
	if err != nil {
		log.Fatalln(err)
	}
	if !match {
		log.Fatalf("IP address %v is not a correct IP address\n", *hostptr)
	}

	port := strconv.Itoa(*portptr)
	host := *hostptr

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	password, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	password = strings.TrimSpace(password)
	fmt.Printf("Generated password = %v\n", password)
}
