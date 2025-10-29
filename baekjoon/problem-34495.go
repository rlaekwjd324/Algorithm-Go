package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
   
	var t int
	fmt.Fscan(reader, &t)
	
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		
		if n % 2 == 0 {
			fmt.Fprintf(writer, "%v%v\n", 2025, strings.Repeat("0", n-4))
			continue
		}
		fmt.Fprintf(writer, "%v%v\n", 42025, strings.Repeat("0", n-5))

	}
}
