package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"math/big"
	"strings"
	"time"

	udt "github.com/vikulin/udt-conn"
)

func main() {
	// utils.SetLogLevel(utils.LogLevelDebug)

	startServer := flag.Bool("s", false, "server")
	startClient := flag.Bool("c", false, "client")
	flag.Parse()

	if *startServer {
		// start the server
		go func() {

			ln, err := udt.Listen("udp", ":8081")
			if err != nil {
				panic(err)
			}

			fmt.Println("Waiting for incoming connection")
			conn, err := ln.Accept()
			if err != nil {
				panic(err)
			}
			fmt.Println("Established connection")

			for {
				message, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					panic(err)
				}
				fmt.Print("Message from client: ", string(message))
				// echo back
				newmessage := strings.ToUpper(message)
				conn.Write([]byte(newmessage + "\n"))
			}
		}()
	}

	if *startClient {
		// run the client
		go func() {
			conn, err := udt.Dial("localhost:8081")
			if err != nil {
				panic(err)
			}

			message := "Ping from client"
			fmt.Fprintf(conn, message+"\n")
			fmt.Printf("Sending message: %s\n", message)
			// listen for reply
			answer, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				panic(err)
			}
			fmt.Print("Message from server: " + answer)
		}()
	}

	time.Sleep(time.Hour)
}
