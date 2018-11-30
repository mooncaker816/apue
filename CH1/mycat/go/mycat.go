package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const BUFFSIZE = 4096

func main() {
	read()
	scan()
}

func scan() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func read() {
	buf := make([]byte, BUFFSIZE)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n > 0 {
			nw, _ := os.Stdout.Write(buf[:n])
			if nw != n {
				log.Fatal("write error")
			}
		}
		if err == io.EOF {
			break
		}
	}
}
