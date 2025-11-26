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

	var a, b string
	var s, v int
	fmt.Fscan(reader, &a, &b, &s, &v)
	
	ah, _ := strconv.Atoi(strings.Split(a, ":")[0])
	am, _ := strconv.Atoi(strings.Split(a, ":")[1])
	as, _ := strconv.Atoi(strings.Split(a, ":")[2])
	bh, _ := strconv.Atoi(strings.Split(b, ":")[0])
	bm, _ := strconv.Atoi(strings.Split(b, ":")[1])
	bs, _ := strconv.Atoi(strings.Split(b, ":")[2])

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

	if v > 0 {
		s = (s*(100-v))/100
	}
	as = as+s
	am = am+as/60
	as = as%60
	ah = ah+am/60
	am = am%60
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
