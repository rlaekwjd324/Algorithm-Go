package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var a, b string
	var s, v int
	fmt.Fscan(reader, &a, &b, &s, &v)
	
	ah := strings.Split(a, ":")[0]
	am := strings.Split(a, ":")[1]
	as := strings.Split(a, ":")[2]
	bh := strings.Split(b, ":")[0]
	bm := strings.Split(b, ":")[1]
	bs := strings.Split(b, ":")[2]

	if ah > bh {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}
	if ah == bh {
		if am > bm {
			fmt.Fprintf(writer, "%v\n", 0)
			return
		}
		if am == bm {
			if as > bs {
				fmt.Fprintf(writer, "%v\n", 0)
				return
			}
		}
	}
	
	fmt.Fprintf(writer, "%v\n", 1)
}
