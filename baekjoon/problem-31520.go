package main

import (
   "bufio"
	"fmt"
	"os"
   "strconv"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var str string
   fmt.Fscan(reader, &str)
   arr := strings.Split(str, "")
   
   pre := 0
   for i := 0; i < len(str); i++ {
      n, _ := strconv.Atoi(arr[i])
      if pre + 1 != n {
         fmt.Fprintln(writer, "-1")
         return
      }
      pre = n
   }
   
   fmt.Fprintf(writer, "%v\n", pre)

}
