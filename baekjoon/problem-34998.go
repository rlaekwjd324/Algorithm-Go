package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	
	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(reader, &c)
		
		sum := 0
		for j := 0; j < c*2+1; j++ {
			var a string
			fmt.Fscan(reader, &a)

			if a == "+" {
				continue
			}
			if a == "!" {
				sum += 10
				continue
			}
			an, _ := strconv.Atoi(a)
			sum += an
		}

		if sum > 9 {
			fmt.Fprintf(writer, "%v\n", "!")
			continue
		}
		fmt.Fprintf(writer, "%v\n", sum)
	}
		
}
