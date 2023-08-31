package main
 
import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	// 1. 2. create and bind (create socket and connect ipaddress:port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println(tcpAddr)
	checkError(err, "tcpAddr")

	// 3. connect server
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "conn")
	fmt.Println("connect server!!")

	// write data to socket
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err, "conn write")
	fmt.Println("conn write")

	// read data from socket
	res := make([]byte, 1024)
	len, err := conn.Read(res)
	checkError(err, "conn read")
	fmt.Println("conn read")
	fmt.Println("response:", string(res[:len]))

	// close connection
	conn.Close()
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s \n", err.Error())
		fmt.Fprintf(os.Stderr, "message: %s \n", msg)
		os.Exit(1)
	}
}
