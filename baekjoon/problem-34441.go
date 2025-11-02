package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)

    var t, w string
    fmt.Fscan(reader, &t, &w)
	var a, p, g int
    fmt.Fscan(reader, &a, &p, &g)

	h, _ := strconv.Atoi(strings.Split(t, ":")[0])
	m, _ := strconv.Atoi(strings.Split(t, ":")[1])
    
	v := h*60+m

	if w == "sun" || w == "sat" {
		v *= 2
	}
	if a == 1 {
		v *= 2
	}
	if p == 1 {
		v *= 3
	}
	if g == 1 {
		v *= 3
	}

	h = v/60
	m = v%60

	if m < 10 {
		fmt.Fprintf(writer, "%v:0%v\n", h, m)
		return
	}
	fmt.Fprintf(writer, "%v:%v\n", h, m)
}
