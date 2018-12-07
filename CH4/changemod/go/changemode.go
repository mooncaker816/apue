package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	prepare()
	for _, file := range files {
		fi, err := os.Stat(file.name)
		if err != nil {
			log.Fatal(err)
		}
		mode := fi.Mode()
		fmt.Printf("before chmod: %o = %s\n", mode, mode)
		// [Min] 注意这里不能直接用 syscall.S_ISGID，Go 的位置与 syscall 的不同
		// [Min] 9位的权限位保持一致，可以用 syscall.S_IXGRP 等
		// err = os.Chmod(file, (mode&^syscall.S_IXGRP)|syscall.S_ISGID)
		if file.name == "foo" {
			err = os.Chmod(file.name, (mode&^syscall.S_IXGRP)|os.ModeSetgid)
		} else {
			err = os.Chmod(file.name, syscall.S_IRUSR|syscall.S_IWUSR|
				syscall.S_IRGRP|syscall.S_IROTH)
		}
		if err != nil {
			log.Fatal(err)
		}
		fi, err = os.Stat(file.name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("after chmod: %s\n", fi.Mode())
	}
}

var (
	files = []struct {
		name  string
		umask int
	}{
		{"foo", 0},
		{"bar", 0066},
	}
)

func prepare() {
	for _, file := range files {
		syscall.Umask(file.umask)
		os.Remove(file.name)
		f, err := os.Create(file.name)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
}
