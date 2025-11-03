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

	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	arr := make([]bool, 26)

    var t string
    fmt.Fscan(reader, &t)
    
	for _, v := range t {
		i := indexOf(string(v), alphabets)
		arr[i] = true
	}
	answer := ""
	for i, v := range arr {
		if !v {
			answer += string(alphabets[i])
		}
	}

	if len(answer) == 0 {
		fmt.Fprintf(writer, "%v\n", "Alphabet Soup!")
		return
	}
	fmt.Fprintf(writer, "%v\n", answer)
}

func indexOf(s, a string) int {
	for i, v := range a {
		if string(v) == s {
			return i
		}
	}
	return -1
}
