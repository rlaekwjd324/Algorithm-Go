package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var t int
   fmt.Fscan(reader, &t)
   
   answer := fmt.Sprintf("%v", t/150)
   t = t%150
   answer = fmt.Sprintf("%v %v", t/30, answer)
   t = t%30
   answer = fmt.Sprintf("%v %v", t/15, answer)
   t = t%15
   answer = fmt.Sprintf("%v %v", t/5, answer)
   t = t%5
   
   fmt.Fprintf(writer, "%v %v\n", t, answer)
}
