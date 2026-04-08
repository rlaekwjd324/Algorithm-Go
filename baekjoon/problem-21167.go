package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var R float64
		var S float64

		_, err := fmt.Fscan(reader, &R, &S)
		if err != nil {
			break // EOF 처리
		}

		// 공식 적용
		V := math.Sqrt((R * (S + 0.16)) / 0.067)

		// 반올림
		result := int(math.Round(V))

		fmt.Fprintln(writer, result)
	}
}
