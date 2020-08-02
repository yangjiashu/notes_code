package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	live := make(chan struct{})
	for {
		go func() {
			if input.Scan() {
				live <- struct{}{}
				echo(conn, input.Text(), 1*time.Second)
			}
		}()
		select {
		case <-live:
		case <-time.After(4 * time.Second):
			fmt.Println("结束连接")
			return
		}
	}
}

func echo(conn net.Conn, words string, delay time.Duration) {
	fmt.Fprintln(conn, strings.ToUpper(words))
	time.Sleep(delay)
	fmt.Fprintln(conn, words)
	time.Sleep(delay)
	fmt.Fprintln(conn, strings.ToLower(words))
}
