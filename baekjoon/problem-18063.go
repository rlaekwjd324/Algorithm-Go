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
   
   var n, c int
   fmt.Fscan(reader, &n, &c)
   
   sumM := 0
   sumS := 0
   for i := 0; i < n; i++ {
      var t string
      fmt.Fscan(reader, &t)

      mm, _ := strconv.Atoi(strings.Split(t, ":")[0])
      ss, _ := strconv.Atoi(strings.Split(t, ":")[1])
      
      sumM += mm
      sumS += ss
   }

   sumM -= 1
   sumS = sumS + 60 - (n - 1) * c
   sumM += sumS / 60
   sumH := sumM / 60
   sumM = sumM % 60
   sumS = sumS % 60

   h := fmt.Sprintf("%v", sumH)
   m := fmt.Sprintf("%v", sumM)
   s := fmt.Sprintf("%v", sumS)
   if sumH < 10 {
      h = fmt.Sprintf("0%v", sumH)
   }
   if sumM < 10 {
      m = fmt.Sprintf("0%v", sumM)
   }
   if sumS < 10 {
      s = fmt.Sprintf("0%v", sumS)
   }
   
   fmt.Fprintf(writer, "%v:%v:%v\n", h, m, s)
}
