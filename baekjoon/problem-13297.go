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

   var n int
   fmt.Fscan(reader, &n)
   
   answer := ""
   for i := 0; i < n; i++ {
       var a string
       fmt.Fscan(reader, &a)
       answer += fmt.Sprintf("%v\n", len(a)) 
   }

   fmt.Fprintln(writer, answer)
}
