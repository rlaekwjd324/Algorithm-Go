package main

import (
    "bufio"
    "fmt"
    "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	max := 0
	min := 1440
	for i := 0; i < 3; i++ {
		var sh, sm, dh, dm int
		fmt.Fscan(reader, &sh, &sm, &dh, &dm)

		v := (dh-sh) * 60 + dm-sm
		if v < 0 {
			v += (24*60)
		}
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	minH := fmt.Sprintf("%v", min/60)
	minM := fmt.Sprintf("%v", min%60)
	if len(minM) == 1 {
		minM = "0"+minM
	}

	maxH := fmt.Sprintf("%v", max/60)
	maxM := fmt.Sprintf("%v", max%60)
	if len(maxM) == 1 {
		maxM = "0"+maxM
	}
	
	fmt.Fprintf(writer, "%v:%v\n", minH, minM)
	fmt.Fprintf(writer, "%v:%v\n", maxH, maxM)
}
