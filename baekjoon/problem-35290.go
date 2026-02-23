package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	if n == 2 || n == 6 || n == 7 || n == 8 {
	    fmt.Fprintln(writer, "Think before submission")
    } else {
    	fmt.Fprintln(writer, "Solve harder problems")
    }
}
