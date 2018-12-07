package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Chdir("/usr/local/mysql")
	if err != nil {
		log.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cwd = %s\n", cwd)
}
