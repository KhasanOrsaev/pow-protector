package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwesterb/go-pow"
)

const difficult = 20

func main() {
	var port int
	flag.IntVar(&port, "port", 8001, "server port")

	flag.Parse()

	rand.Seed(time.Now().Unix())
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{Port: port})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
Scan:
	for scanner.Scan() {
		message := scanner.Text()

		switch strings.TrimSpace(message) {
		case "exit":
			break Scan
		case "get":
			req := pow.NewRequest(difficult, []byte(strconv.Itoa(rand.Int())))
			_, err := fmt.Fprintf(conn, req+"\n")
			if err != nil {
				fmt.Println(err)
				break Scan
			}
			scanner.Scan()
			proof := scanner.Text()
			ok, err := pow.Check(req, proof, []byte("some bound data"))
			if err != nil {
				fmt.Println(err)
			}
			if ok {
				_, err := fmt.Fprintf(conn, "Word of wizdom \n")
				if err != nil {
					fmt.Println(err)
					break Scan
				}
			}
		default:
			_, err := fmt.Fprintf(conn, "Unknown command \n")
			if err != nil {
				fmt.Println(err)
				break Scan
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
