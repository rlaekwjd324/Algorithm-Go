package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

    var n int
    fmt.Fscan(reader, &n)
	
	for i := 0; i <= n; i++ {
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		if len(s) >= 10 {
			if strings.Contains(s[:10], "Simon says") {
				fmt.Fprintf(writer, "%v\n", s[10:])
			}
		}
	}
}
