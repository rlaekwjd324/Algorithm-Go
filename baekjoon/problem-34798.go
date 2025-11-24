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

	var b, a string
	fmt.Fscan(reader, &b, &a)

	bh := strings.Split(b, ":")[0]
	bm := strings.Split(b, ":")[1]
	ah := strings.Split(a, ":")[0]
	am := strings.Split(a, ":")[1]

	if ah > bh {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}
	if ah == bh && am > bm {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}
	
	fmt.Fprintf(writer, "%v\n", "NO")
}
