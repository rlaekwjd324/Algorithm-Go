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

	var t1, v1, t2, v2 int
	fmt.Fscan(reader, &t1, &v1, &t2, &v2)
	
	if t2 < 0 && v2 >= 10 {
		fmt.Fprintf(writer, "%v\n", "A storm warning for tomorrow! Be careful and stay home if possible!")
		return
	}
	if t2 < t1 {
		fmt.Fprintf(writer, "%v\n", "MCHS warns! Low temperature is expected tomorrow.")
		return
	}
	if v2 > v1 {
		fmt.Fprintf(writer, "%v\n", "MCHS warns! Strong wind is expected tomorrow.")

	}
	fmt.Fprintf(writer, "%v\n", "No message")
}
