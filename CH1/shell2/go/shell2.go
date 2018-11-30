package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

func main() {
	fmt.Println("% ")
	ch := make(chan os.Signal)
	go func() {
		for {
			<-ch
			fmt.Print("interrupt\n%")
		}
	}()
	signal.Notify(ch, os.Interrupt)
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
