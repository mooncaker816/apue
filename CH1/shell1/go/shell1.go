package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("% ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		args := strings.SplitN(s.Text(), " ", 2)
		cmd := exec.Command(args[0], args[1:]...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(out))
		fmt.Println("% ")
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
