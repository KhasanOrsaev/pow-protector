package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/bwesterb/go-pow"
)

var data = []byte("some bound data")

func main() {
	var port int
	var host string
	flag.IntVar(&port, "port", 8001, "server port")
	flag.StringVar(&host, "host", "127.0.0.1", "server host")
	flag.Parse()

	// Connect to host
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		// stdin reader
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		switch strings.Trim(command, "\n") {
		case "get", "exit":
		default:
			fmt.Println("Unknown command. Try 'get' or 'exit'")
			continue
		}

		// send
		_, err = fmt.Fprintf(conn, command)
		if err != nil {
			fmt.Println(err)
			break
		}

		// listen response
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		proof, err := pow.Fulfil(message, data)
		if err != nil {
			fmt.Println(err)
			break
		}
		// send proof
		_, err = fmt.Fprintf(conn, proof+"\n")
		if err != nil {
			fmt.Println(err)
			break
		}

		// listen response
		message, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		if message != "forbidden" {
			fmt.Print("Word of wizdom: ", message)
		}
	}
}
