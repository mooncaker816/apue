package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("hello world from process ID %d\n", os.Getpid())
}
