package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	hashMap := map[int]int{
		232:  2017,
		88:   2018,
		754:  2019,
		29:   2020,
		28:   2021,
		1015: 2022,
		1295: 2023,
		1073: 2024,
		348:  2025,
	}

	var T int
	fmt.Fscan(in, &T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Fscan(in, &N)

		product := 1

		for i := 0; i < N; i++ {
			var s string
			fmt.Fscan(in, &s)
			product *= len(s)
		}

		if year, ok := hashMap[product]; ok {
			fmt.Fprintln(out, year)
		} else {
			fmt.Fprintln(out, "Goodbye, ChAOS!")
		}
	}
}
