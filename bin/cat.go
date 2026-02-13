package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader, setNu bool) {
	fmt.Println(setNu)
	lineIndex := 1
	for {
		buf, err := r.ReadBytes('\n')
		line := fmt.Sprintf("%s", string(buf))
		if setNu {
			line = fmt.Sprintf("%d %s", lineIndex, string(buf))
		}
		//fmt.Println(line)
		lineIndex += 1
		fmt.Fprintf(os.Stdout, line)
		if err == io.EOF {
			break
		}
	}
	return
}
func main() {
	setNu := flag.Bool("n", false, "每一行头部加入一个行号")
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin), false)
	}
	for i := 0; i < flag.NArg(); i++ {
		arg := flag.Arg(i)
		fmt.Println()
		if arg == "-n" {
			continue
		}
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f), *setNu)
		f.Close()
	}
}
