package main

import (
	"fmt"
	"log"
	"os"
)

var (
	buf1 = "abcdefghij"
	buf2 = "ABCDEFGHIJ"
)

func main() {
	f, err := os.Create("hole.file")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	n, err := f.Write([]byte(buf1))
	if err != nil || n != 10 {
		log.Fatal("write error")
	}
	p, err := f.Seek(16384, os.SEEK_SET)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	n, err = f.Write([]byte(buf2))
	if err != nil || n != 10 {
		log.Fatal("write error")
	}
	nohole()
}

func nohole() {
	f, err := os.Create("nohole.file")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for i := 0; i < 16394; i++ {
		_, err := f.Write([]byte{'1'})
		if err != nil {
			panic(err)
		}
	}
}
