package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanBytes)
	for s.Scan() {
		fmt.Print(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
