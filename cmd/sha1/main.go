package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
)

func main() {
	var data string

	if len(os.Args) > 1 {
		data = strings.Join(os.Args[1:], " ")
	} else {
		// no data so trying stdin
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			data = s.Text()
		}
	}

	h := sha1.New()
	h.Write([]byte(data))

	fmt.Printf("%x\n", h.Sum(nil))
}
