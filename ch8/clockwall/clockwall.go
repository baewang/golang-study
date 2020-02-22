package main

import (
	flag2 "flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	newYork := setConn("tcp", *flag2.String("NewYork", "localhost:8010", "NewYork Server Port"))
	mustCopy(os.Stdout, newYork)

	london := setConn("tcp", *flag2.String("London", "localhost:8020", "London Server Port"))
	mustCopy(os.Stdout, london)

	tokyo := setConn("tcp", *flag2.String("Tokyo", "localhost:8030", "Tokyo Server Port"))
	mustCopy(os.Stdout, tokyo)
	defer newYork.Close()
	defer london.Close()
	defer tokyo.Close()
}

func setConn(network, address string) net.Conn {
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
