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

   var n int
   var s string
   fmt.Fscan(reader, &n, &s)
   
   sp := ".#:|"
   sum := 0
   for _, v := range s {
      if strings.Contains(sp, string(v)) {
         idx := indexOf(s, string(v))
         t, _ := strconv.Atoi(s[:idx])
         sum += t
         s = s[idx+1:]
      }
   }
   t, _ := strconv.Atoi(s)
   sum += t
   fmt.Fprintf(writer, "%v\n", sum)
}

func indexOf(s, t string) int {
   for i, v := range s {
      if string(v) == t {
         return i
      }
   }
   return -1
}
