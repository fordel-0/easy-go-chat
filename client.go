package main


import (
	"fmt"
	"net"
	"bufio"
	"os"
)


var inpReader = bufio.NewReader(os.Stdin)


func main() {
	fmt.Print("Server address: ")

	servAddr, _, err := inpReader.ReadLine()

	conn, err := net.Dial("tcp", string(servAddr))
	p(err)
	defer conn.Close()

	go feedback(conn)

	for {
		_, err := conn.Write(input())
		p(err)
	}
}


func input() ([]byte) {
	s, _, err := inpReader.ReadLine()
	p(err)

	fmt.Print("\033[1A\033[2K\r")
	return s
}


func feedback(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		m, _, err := reader.ReadLine()
		p(err)
		fmt.Print(string(m) + "\n")
	}
}


func p(err error) {
	if err != nil {panic(err)}
}
