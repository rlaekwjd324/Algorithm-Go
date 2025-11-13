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

    var n string
    fmt.Fscan(reader, &n)
	
	if string(n[0]) == "-" {
		n = n[1:]

		if len(n) < 6 {
			ni, _ := strconv.Atoi(n)
			if ni <= 32768 {
				fmt.Fprintf(writer, "%v\n", "short")
				return
			}
		}
	
		if len(n) < 11 {
			ni, _ := strconv.Atoi(n)
			if ni <= 2147483648 {
				fmt.Fprintf(writer, "%v\n", "int")
				return
			}
		}
	
		fmt.Fprintf(writer, "%v\n", "long long")
		return
	}

	if len(n) < 6 {
		ni, _ := strconv.Atoi(n)
		if ni < 32768 {
			fmt.Fprintf(writer, "%v\n", "short")
			return
		}
	}

	if len(n) < 11 {
		ni, _ := strconv.Atoi(n)
		if ni < 2147483648 {
			fmt.Fprintf(writer, "%v\n", "int")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "long long")
}
