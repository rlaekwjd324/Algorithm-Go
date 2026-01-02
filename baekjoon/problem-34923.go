package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var n, m string
	fmt.Fscan(reader, &n, &m)

	bH, _ := strconv.Atoi(strings.Split(n, ":")[0])
	bM, _ := strconv.Atoi(strings.Split(n, ":")[1])
	aH, _ := strconv.Atoi(strings.Split(m, ":")[0])
	aM, _ := strconv.Atoi(strings.Split(m, ":")[1])
	
	rM := aM-bM
	rH := aH-bH
	if rM < 0 {
		rH--
		rM += 60
	}
	if rH < 0 {
		rH += 12
	}
	r := (rM + rH * 60) * 6

	lM := bM-aM
	lH := bH-aH
	if lM < 0 {
		lH--
		lM += 60
	}
	if lH < 0 {
		lH += 12
	}
	l := (lM + lH * 60) * 6

	if r < l {
		fmt.Fprintf(writer, "%v\n", r)
		return
	}
	fmt.Fprintf(writer, "%v\n", l)
}
