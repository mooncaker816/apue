package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

const (
	X_OK = 1 << iota
	W_OK
	R_OK
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: access <pathname>")
	}
	err := syscall.Access(os.Args[1], R_OK)
	if err != nil {
		log.Fatalf("access error for %s: %v\n", os.Args[1], err)
	}
	fmt.Println("read access OK")
	_, err = os.OpenFile(os.Args[1], os.O_RDONLY, 0)
	if err != nil {
		log.Fatalf("open error for %s: %v\n", os.Args[1], err)
	}
	fmt.Println("open for reading OK")
}
