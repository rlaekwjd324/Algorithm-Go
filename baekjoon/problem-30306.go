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

	var n int
	fmt.Fscan(reader, &n)
	
	fArr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &fArr[i])
	}
	sArr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &sArr[i])
	}

	fw, sw := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if fArr[i] > sArr[j] {
				fw++
				continue
			}
			if fArr[i] < sArr[j] {
				sw++
			}
		}
	}
	
	if fw > sw {
		fmt.Fprintf(writer, "%v\n", "first")
		return
	}
	if fw < sw {
		fmt.Fprintf(writer, "%v\n", "second")
		return
	}
	fmt.Fprintf(writer, "%v\n", "tie")
}
