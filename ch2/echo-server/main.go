package main

import (
	"log"
	"net"
	"io"
)

//echo is a handler function that simply echos received data.

func echo(conn net.Conn)  {
	defer conn.Close()

	//Create a buffer to store received data.
	b := make([]byte, 512)
	for {
		//Receive data via conn.Read into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Printf("Unexpected error")
			break
		}
		log.Printf("Recieved %d bytes: %s\n", size, string(b))

		//Send data via conn.Write.
		log.Printf("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main()  {
	//Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for{
		//Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received Connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		//Handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}

}