package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

type Interval struct {
	start int
	end   int
}

func toMinutes(h, m int) int {
	return h*60 + m
}

func main() {
	var h, m int
	fmt.Fscan(reader, &h, &m)

	t := toMinutes(h, m)

	intervals := []Interval{
		{toMinutes(6, 30), toMinutes(9, 0)},
		{toMinutes(9, 50), toMinutes(10, 0)},
		{toMinutes(10, 50), toMinutes(11, 0)},
		{toMinutes(11, 50), toMinutes(12, 0)},
		{toMinutes(12, 50), toMinutes(13, 50)},
		{toMinutes(14, 40), toMinutes(14, 50)},
		{toMinutes(15, 40), toMinutes(15, 50)},
		{toMinutes(16, 40), toMinutes(22, 50)},
	}

	for _, iv := range intervals {
		if t >= iv.start && t <= iv.end {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
