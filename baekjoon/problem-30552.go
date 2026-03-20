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
	
	pre := ""
	c := 0
	for i := 0; i < n; i++ {
		var name string
		fmt.Fscan(reader, &name)
		
		if name == "Present!" {
			pre = ""
			c--
			continue
		}
		if pre != "" {
			fmt.Fprintf(writer, "%v\n", pre)
		}
		pre = name
		c++
	}

	if c == 0 {
		fmt.Fprintf(writer, "%v\n", "No Absences")
		return
	}
	if pre != "" {
		fmt.Fprintf(writer, "%v\n", pre)
	}
}
