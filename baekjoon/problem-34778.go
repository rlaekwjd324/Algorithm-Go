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

	arr := make([]bool, 4)
	for i := 0; i < 10; i++ {
		var a int
		fmt.Fscan(reader, &a)
		arr[a-1] = true
	}

	sum := 0
	for i := 0; i < 4; i++ {
		if !arr[i] {
			sum++
		}
	}
	
	fmt.Fprintf(writer, "%v\n", sum)
}
