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

	for {
		var n string
		fmt.Fscan(reader, &n)
		
		switch n {
		case "animal":
			fmt.Fprintf(writer, "%v\n", "Panthera tigris")
		case "tree":
			fmt.Fprintf(writer, "%v\n", "Pinus densiflora")
		case "flower":
			fmt.Fprintf(writer, "%v\n", "Forsythia koreana")
		case "end":
			return
		}
	}
	
}
