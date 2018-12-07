package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var c1, c2 counts

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage:  ftw  <starting-pathname>")
	}
	// similar to c program
	err := walkPath(os.Args[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Println(c1)

	// exsiting filepath.Walk
	err = filepath.Walk(os.Args[1], walkCalc)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(c2)
}

func walkPath(path string) error {
	fi, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("stat error for %s: %v", path, err)
	}
	if !fi.IsDir() {
		return c1.calc(fi.Mode())
	}
	c1.ndir++
	dir, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("can't open directory %s", path)
	}
	defer dir.Close()

	names, err := dir.Readdirnames(0)
	if err != nil {
		return fmt.Errorf("can't read directory %s", path)
	}
	for _, name := range names {
		newPath := filepath.Join(path, name)
		err := walkPath(newPath)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

type counts struct {
	nreg, ndir, nblk, nchr, nfifo, nslink, nsock, ntot float32
}

// case syscall.S_IFBLK:
// 	fs.mode |= ModeDevice
// case syscall.S_IFCHR:
// 	fs.mode |= ModeDevice | ModeCharDevice
// case syscall.S_IFDIR:
// 	fs.mode |= ModeDir
// case syscall.S_IFIFO:
// 	fs.mode |= ModeNamedPipe
// case syscall.S_IFLNK:
// 	fs.mode |= ModeSymlink
// case syscall.S_IFREG:
// 	// nothing to do
// case syscall.S_IFSOCK:
// 	fs.mode |= ModeSocket
func (c *counts) calc(mode os.FileMode) error {
	switch {
	case mode.IsRegular():
		c.nreg++
	case mode&os.ModeCharDevice != 0 && mode&os.ModeDevice != 0:
		c.nchr++
	case mode&os.ModeDevice != 0:
		c.nblk++
	case mode&os.ModeNamedPipe != 0:
		c.nfifo++
	case mode&os.ModeSymlink != 0:
		c.nslink++
	case mode&os.ModeSocket != 0:
		c.nsock++
	}
	return nil
}

func (c counts) String() string {
	var b strings.Builder
	ntot := c.nreg + c.ndir + c.nblk + c.nchr + c.nfifo + c.nslink + c.nsock
	if ntot == 0 {
		ntot = 1
	}
	b.WriteString(fmt.Sprintf("total = %d\n", int(ntot)))
	b.WriteString(fmt.Sprintf("regular files  = %7d, %5.2f %%\n", int(c.nreg), c.nreg*100.0/ntot))
	b.WriteString(fmt.Sprintf("directories  = %7d, %5.2f %%\n", int(c.ndir), c.ndir*100.0/ntot))
	b.WriteString(fmt.Sprintf("block special  = %7d, %5.2f %%\n", int(c.nblk), c.nblk*100.0/ntot))
	b.WriteString(fmt.Sprintf("char special  = %7d, %5.2f %%\n", int(c.nchr), c.nchr*100.0/ntot))
	b.WriteString(fmt.Sprintf("FIFOs  = %7d, %5.2f %%\n", int(c.nfifo), c.nfifo*100.0/ntot))
	b.WriteString(fmt.Sprintf("symbolic links  = %7d, %5.2f %%\n", int(c.nslink), c.nslink*100.0/ntot))
	b.WriteString(fmt.Sprintf("sockets  = %7d, %5.2f %%\n", int(c.nsock), c.nsock*100.0/ntot))
	return b.String()
}

func walkCalc(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		c2.ndir++
		return nil
	}
	return c2.calc(info.Mode())
}
