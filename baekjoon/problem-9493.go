package main

import (
    "bufio"
    "fmt"
    "os"
	"math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	for {
		var m, a, b float64
		fmt.Fscan(reader, &m, &a, &b)

		if m == 0 && a == 0 && b == 0 {
			return
		}

		// akm/h akm/3600s m : s = a : 3600 m*3600=a*s s = m*3600/a
		va := m*3600/a
		vb := m*3600/b
		v := int(math.Round(va-vb))
		if v < 0 {
			v = v*-1
		}
		h := fmt.Sprintf("%v", v/3600)
		v = v%3600
		mm := fmt.Sprintf("%v", v/60)
		if len(mm) == 1 {
			mm = "0"+mm
		}
		ss := fmt.Sprintf("%v", v%60)
		if len(ss) == 1 {
			ss = "0"+ss
		}
		fmt.Fprintf(writer, "%v:%v:%v\n", h, mm, ss)
	}
}
