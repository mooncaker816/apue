package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: fileflags <descriptor#>")
	}
	fd, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	flag, err := fcntl(fd, syscall.F_GETFL, 0)
	if err != nil {
		log.Fatal(err)
	}
	switch flag & syscall.O_ACCMODE {
	case syscall.O_RDONLY:
		fmt.Print("read only")
	case syscall.O_WRONLY:
		fmt.Print("write only")
	case syscall.O_RDWR:
		fmt.Print("read write")
	default:
		log.Fatal("unknow access mode")
	}
	if flag&syscall.O_APPEND != 0 {
		fmt.Print(", append")
	}
	if flag&syscall.O_NONBLOCK != 0 {
		fmt.Print(", nonblocking")
	}
	if flag&syscall.O_SYNC != 0 {
		fmt.Print(", synchronous writes")
	}
	fmt.Println()
}

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
