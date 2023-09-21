package main


import (
	"fmt"
	"net"
	"bufio"
	"os"
)


const (
	servAddr = ":44821"
)


func main() {
	conn, err := net.Dial("tcp", servAddr)
	p(err)
	defer conn.Close()

	/*l, err := net.Listen("tcp", ":")
    p(err)
    defer l.Close()

	conn.Write()*/

	go feedback(conn)

	for {
		conn.Write(input())
		p(err)
	}
}


func input() ([]byte) {
	reader := bufio.NewReader(os.Stdin)

	s, err := reader.ReadString('\n')
	p(err)

	fmt.Print("\033[1A\033[2K\r")
	return []byte(s)
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
