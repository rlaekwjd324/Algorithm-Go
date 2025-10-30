
package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
   
	var s string
	fmt.Fscan(reader, &s)
	
	x := int(s[0]-'a')
	y := int(s[1]-'1')

	moves := [8][2]int{
		{2, 1},   // →→↑
		{2, -1},  // →→↓
		{-2, 1},  // ←←↑
		{-2, -1}, // ←←↓
		{1, 2},   // ↑↑→
		{-1, 2},  // ↑↑←
		{1, -2},  // ↓↓→
		{-1, -2}, // ↓↓←
	}

	var arr []string
	for i := 0; i < 8; i++ {
		nx := x+moves[i][0]
		ny := y+moves[i][1]
		if nx >= 0 && nx < 8 && ny >= 0 && ny < 8 {
			arr = append(arr, string(rune(nx+'a'))+string(rune(ny+'1')))
		}
	}

	sort.Strings(arr)

	for _, v := range arr {
		fmt.Fprintf(writer, "%v\n", v)
	}
}
