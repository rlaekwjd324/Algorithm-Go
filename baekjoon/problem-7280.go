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

	sArr := make([]bool, 13)
	bArr := make([]bool, 13)
	vArr := make([]bool, 13)
	kArr := make([]bool, 13)
	s, b, v, k := 0, 0, 0, 0
	for i := 0; i < 51; i++ {
		var t string
		var n int
		fmt.Fscan(reader, &t, &n)

		switch t {
		case "S":
			sArr[n-1] = true
			s++
		case "B":
			bArr[n-1] = true
			b++
		case "V":
			vArr[n-1] = true
			v++
		case "K":
			kArr[n-1] = true
			k++
		}
	}
	
	if s == 12 {
		for i := 0; i < 13; i++ {
			if !sArr[i] {
				fmt.Fprintf(writer, "%v %v\n", "S", i+1)
				return
			}
		}
	}
	if b == 12 {
		for i := 0; i < 13; i++ {
			if !bArr[i] {
				fmt.Fprintf(writer, "%v %v\n", "B", i+1)
				return
			}
		}
	}
	if v == 12 {
		for i := 0; i < 13; i++ {
			if !vArr[i] {
				fmt.Fprintf(writer, "%v %v\n", "V", i+1)
				return
			}
		}
	}
	for i := 0; i < 13; i++ {
		if !kArr[i] {
			fmt.Fprintf(writer, "%v %v\n", "K", i+1)
			return
		}
	}
}
