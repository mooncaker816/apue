package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

// [Min] umask 可以看成是在你给定的文件权限基础上扣除 umask 值（8进制）对应的权限

func main() {
	// [Min] 假定我想建一个777的 foo 文件，但是由于默认的 umask 存在，
	// [Min] 会从777中扣除 umask 对应的权限，得到的是：
	// [Min] -rwxr-xr-x  1 mingle  staff        0 12  7 09:19 foo
	// [Min] 说明扣除了组，其他的写权限，即 umask 为0022
	_, err := os.OpenFile("foo", os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	oldmask := syscall.Umask(0)
	fmt.Printf("%o\n", oldmask) // [Min] oldmask 为22，证明了上面的假设是对的
	// [Min] 由于将 umask 置为了0，所以文件的权限就会按照给定的值来建，
	// [Min] foo1 就会是777，
	// [Min] -rwxrwxrwx  1 mingle  staff        0 12  7 09:19 foo1
	_, err = os.OpenFile("foo1", os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	// [Min] 再将 umask 置为 066，即扣除组，其他的读，写
	oldmask = syscall.Umask(0066)
	fmt.Printf("%o\n", oldmask)
	_, err = os.OpenFile("bar", os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// [Min] 最终 bar 符合预期
	// [Min] -rw-------  1 mingle  staff        0 12  7 09:19 bar
}
