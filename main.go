package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number")
		return
	}

	PORT := ":" + arguments[1]
	fmt.Println("here aroundment it ", PORT)
	l, err := net.Listen("tcp4", PORT) //creates a listerner
	if err != nil {
		log.Fatal(err)
		return
	}
	defer l.Close() // runs l.close when main function ends

	for {
		c, err := l.Accept() //three way handshake // c-> connecttion object of type net.Conn basically a socket (private phone line between client and server)

		if err != nil {
			log.Fatal(err)
			return
		}

		go handleConnection(c) // go routine -> server can go back to handling more clients, and can handle connections concurrently

	}

}

func handleConnection(c net.Conn) {
	fmt.Println("Serving ", c.RemoteAddr().String())
	buffer := make([]byte, 1024)
	defer c.Close() // four way handshake // once done talking to client close the connection

	for {
		n, err := c.Read(buffer)
		if err != nil {
			fmt.Println("There was osme error ", err)
			return
		}
		data := buffer[:n]
		fmt.Println(string(data))
		_, errw := c.Write(buffer[:n])
		if errw != nil {

			fmt.Println("There was some erreor writing ", err) // shown when client terminal the session
			return
		}
	}

}
