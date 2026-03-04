package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, x, m int
	fmt.Fscan(reader, &n, &x, &m)

	crate := make([]int, n+1)

	for i := 0; i < m; i++ {
		var r int
		fmt.Fscan(reader, &r)
		crate[r]++
	}

	// 시계 방향 비용
	clockwise := 0
	cur := 1
	for cur != x {
		clockwise += crate[cur]
		cur++
		if cur > n {
			cur = 1
		}
	}

	// 반시계 방향 비용
	counter := 0
	cur = 1
	for cur != x {
		prev := cur - 1
		if prev == 0 {
			prev = n
		}
		counter += crate[prev]
		cur = prev
	}

	if clockwise < counter {
		fmt.Fprintln(writer, clockwise)
	} else {
		fmt.Fprintln(writer, counter)
	}
}
