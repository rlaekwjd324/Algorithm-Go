package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var n, m, s int
	fmt.Fscan(reader, &n, &m)
	
	nArr := make([]string, n)
	mArr := make([]string, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nArr[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &mArr[i])
	}
	fmt.Fscan(reader, &s)
	
	c := 86400
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			v := getSec(nArr[i], mArr[j])
			if v >= s && c > v {
				c = v
			}
		}
	}

	if c == 86400 {
		fmt.Fprintf(writer, "%v\n", -1)
		return
	}
	fmt.Fprintf(writer, "%v\n", c)
}

func getSec(n, m string) int {
	nh, _ := strconv.Atoi(strings.Split(n, ":")[0])
	nm, _ := strconv.Atoi(strings.Split(n, ":")[1])
	ns, _ := strconv.Atoi(strings.Split(n, ":")[2])
	mh, _ := strconv.Atoi(strings.Split(m, ":")[0])
	mm, _ := strconv.Atoi(strings.Split(m, ":")[1])
	ms, _ := strconv.Atoi(strings.Split(m, ":")[2])
	
	nv := nh*3600+nm*60+ns
	mv := mh*3600+mm*60+ms

	return mv-nv
}
