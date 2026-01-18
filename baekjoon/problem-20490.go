package main

import (
    "bufio"
    "fmt"
    "os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	fArr, sArr := make([]int, 3), make([]int, 3)
	for i := 0; i < 6; i++ {
		if i < 3 {
			fmt.Fscan(reader, &fArr[i])
			continue
		}
		fmt.Fscan(reader, &sArr[i-3])
	}

	sort.Ints(fArr)
	sort.Ints(sArr)
	
	v := fArr[0]+fArr[1]+sArr[0]+sArr[1]
	d := fArr[2]-sArr[2]
	if d < 0 {
		d = d * -1
	}
	
	fmt.Fprintf(writer, "%v\n", v+d)
}
