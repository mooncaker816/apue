package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Printf("%s: ", arg)
		fi, err := os.Lstat(arg)
		if err != nil {
			log.Fatal(err)
		}
		mode := fi.Mode()
		switch {
		case mode.IsRegular():
			fmt.Println("regular")
		case mode.IsDir():
			fmt.Println("directory")
		case mode&os.ModeCharDevice != 0 && mode&os.ModeDevice != 0:
			fmt.Println("character special")
		case mode&os.ModeDevice != 0:
			fmt.Println("block special")
		case mode&os.ModeNamedPipe != 0:
			fmt.Println("fifo")
		case mode&os.ModeSymlink != 0:
			fmt.Println("symbolic link")
		case mode&os.ModeSocket != 0:
			fmt.Println("socket")
		default:
			fmt.Println("** unknown mode **")
		}
	}
}

// [Min] on my linux machine
// ./filetype /etc/passwd /etc /dev/log /dev/tty /run/systemd/inaccessible/fifo /dev/sr0 /dev/cdrom
// /etc/passwd: regular
// /etc: directory
// /dev/log: symbolic link
// /dev/tty: character special
// /run/systemd/inaccessible/fifo: fifo
// /dev/sr0: block special
// /dev/cdrom: symbolic link
