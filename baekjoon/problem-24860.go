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

	var vk, jk, vl, jl, vh, dh, jh int
	fmt.Fscan(reader, &vk, &jk, &vl, &jl, &vh, &dh, &jh)

	v := (vk*jk+vl*jl)*(vh*dh*jh)
	fmt.Fprintf(writer, "%v\n", v)
}
