package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
)

func main() {
	fmt.Printf("uid = %d, gid = %d\n", os.Getuid(), os.Getgid())
	userpkg()
}

func userpkg() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("uid = %s, gid = %s\n", u.Uid, u.Gid)
	for i := 0; i < 1000; i++ {
		u, err := user.LookupId(strconv.Itoa(i))
		if err != nil {
			//log.Printf("uid - %d not exist!\n", i)
			continue
		}
		fmt.Printf("%+v\n", u)
	}
}
