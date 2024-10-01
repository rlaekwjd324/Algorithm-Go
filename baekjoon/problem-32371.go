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

	var arr = []string{}
	for i := 0; i < 4; i++ {
		var temp string
		fmt.Fscanln(reader, &temp)
		arr = append(arr, temp)
	}

	var key = [4][10]bool{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			key[i][j] = false
		}
	}
	var str string
	fmt.Fscanln(reader, &str)
	strArr := strings.Split(str, "")
	for i := 0; i < len(str); i++ {
		for j := 0; j < 4; j++ {
			if strings.Contains(arr[j], strArr[i]) {
				key[j][strings.Index(arr[j], strArr[i])] = true
				break
			}
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			if key[i][j] && key[i+1][j] && key[i+2][j] && key[i][j+1] && key[i+1][j+1] && key[i+2][j+1] && key[i][j+2] && key[i+1][j+2] && key[i+2][j+2] {
				answer := strings.Split(arr[i+1], "")
				fmt.Fprintln(writer, answer[j+1])
				return
			}
		}
	}
}
