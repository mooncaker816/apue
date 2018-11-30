package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: ls directory_name")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fis, err := f.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fis {
		fmt.Println(fi.Name())
	}
}
