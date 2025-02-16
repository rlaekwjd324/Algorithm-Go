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
   arrC := make([]string, n)
   arrW := make([]string, n)
   for i := 0; i < n; i++ {
      var c, w string
      fmt.Fscan(reader, &c, &w)
      
      arrC[i] = c
      arrW[i] = w
   }
   var p string
   fmt.Fscan(reader, &p)
   a := ""
   for i := 0; i <= len(p)-4; i=i+4 {
      s := string(p[i])+string(p[i+1])+string(p[i+2])+string(p[i+3])
      index := indexOf(s, arrW)
      if index == -1 {
         a += "?"
         continue
      } 
      a += string(arrC[index])
   }
   
   fmt.Fprint(writer, a)
}

func indexOf(t string, arr []string) int {
   for i, v := range arr {
      if t == v {
         return i
      }
   }
   return -1
}
