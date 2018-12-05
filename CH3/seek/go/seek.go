package main

import (
	"fmt"
	"os"
)

func main() {
	offset, err := os.Stdin.Seek(0, os.SEEK_CUR)
	if err != nil {
		fmt.Printf("cannot seek %v\n", err)
		return
	}
	fmt.Printf("seek ok: %d\n", offset)
}
