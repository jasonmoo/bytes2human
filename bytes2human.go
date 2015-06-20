package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	divisor   float64 = 1024
	read_size         = 4 << 10 // 4kb read buf
)

var (
	suffixes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	help     = flag.Bool("?", false, "display help")
)

func main() {

	var r *regexp.Regexp

	if len(os.Args) == 2 {
		r = regexp.MustCompile(os.Args[1])
	} else {
		r = regexp.MustCompile(`\d+`)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Err() != nil {
			fmt.Println(s.Err())
			break
		}
		os.Stdout.WriteString(r.ReplaceAllStringFunc(s.Text(), func(input string) string {

			// fmt.Println("input", input)

			n, err := strconv.ParseFloat(input, 64)
			if err != nil || n < 1024 {
				return input + suffixes[0]
			}

			var negator float64 = 1

			if n < 0 {
				negator = -negator
				n *= negator
			}

			var i int
			for n >= divisor {
				i, n = i+1, n/divisor
			}

			return strconv.FormatFloat(n*negator, 'f', 0, 64) + suffixes[i]

		}))

	}

}
