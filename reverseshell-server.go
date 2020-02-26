package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
) 
func main() {
	listener, _ := net.Listen("tcp","127.0.0.1:1337")
	CLIReader := bufio.NewReader(os.Stdin)
	for {
		conn,_ := listener.Accept()

		for {
			fmt.Printf("Execute > ")
			cmd,_ := CLIReader.ReadString('\n')
			conn.Write([]byte(cmd))
			
			Session(conn)
		}
	}
}

func Session(conn net.Conn) {
	R := bufio.NewReader(conn)	
	buf := make([]byte, 1024)
	result,err := R.Read(buf)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(string(buf[:result]))
	}
}