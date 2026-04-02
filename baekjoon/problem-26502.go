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

	var t int
	fmt.Fscan(reader, &t)
	reader.ReadString('\n')
	
	var sb strings.Builder
	for i := 0; i < t; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimRight(str, "\n")
		
		for _, v := range str {
			switch v {
			case 'a':
				sb.WriteByte('e')
			case 'e':
				sb.WriteByte('i')
			case 'i':
				sb.WriteByte('o')
			case 'o':
				sb.WriteByte('u')
			case 'u':
				sb.WriteByte('y')
			case 'y':
				sb.WriteByte('a')
			case 'A':
				sb.WriteByte('E')
			case 'E':
				sb.WriteByte('I')
			case 'I':
				sb.WriteByte('O')
			case 'O':
				sb.WriteByte('U')
			case 'U':
				sb.WriteByte('Y')
			case 'Y':
				sb.WriteByte('A')
			default:
				sb.WriteRune(v)
			}
		}
		sb.WriteByte('\n')
	}

	fmt.Fprintf(writer, "%v", sb.String())
}
