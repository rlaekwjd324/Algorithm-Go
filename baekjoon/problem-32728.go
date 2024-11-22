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

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	var balls string
	fmt.Fscanln(reader, &balls)
	ballArr := strings.Split(balls, "")

	boxs := []string{"", "", ""}
	for i := 0; i < n; i++ {
		switch ballArr[i] {
		case "s":
			if len(boxs[0]) < k {
				boxs[0] += ballArr[i]
			} else if len(boxs[1]) < k {
				boxs[1] += ballArr[i]
			} else {
				boxs[2] += ballArr[i]
			}
		case "r":
			if len(boxs[1]) < k {
				boxs[1] += ballArr[i]
			} else if len(boxs[2]) < k {
				boxs[2] += ballArr[i]
			} else {
				boxs[0] += ballArr[i]
			}
		case "p":
			if len(boxs[2]) < k {
				boxs[2] += ballArr[i]
			} else if len(boxs[0]) < k {
				boxs[0] += ballArr[i]
			} else {
				boxs[1] += ballArr[i]
			}
		}
	}
	fmt.Fprintf(writer, "%v\n%v\n%v\n", boxs[0], boxs[1], boxs[2])
}
