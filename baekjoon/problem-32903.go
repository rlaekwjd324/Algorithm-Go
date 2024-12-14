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
	fmt.Fscanln(reader, &n)

	al := make([]string, 26)
	for i := 0; i < n; i++ {
		var app string
		fmt.Fscanln(reader, &app)

		first := strings.Split(app, "")[0]
		al[first[0]-97] = first
	}

	for i := 0; i < 26; i++ {
		if al[i] == "" {
			fmt.Fprint(writer, ".")
		} else {
			fmt.Fprint(writer, al[i])
		}
		if i%6 == 5 {
			fmt.Fprint(writer, "\n")
		}
	}
}
