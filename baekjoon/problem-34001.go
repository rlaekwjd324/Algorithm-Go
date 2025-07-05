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

   var t int
   fmt.Fscan(reader, &t)

   aArr := [][]int{
      {200,	210,	220},
      {210,	220,	225},
      {220,	225,	230},
      {225,	230,	235},
      {230,	235,	245},
      {235,	245,	250},
      {260,	265,	270},
      {265,	270,	275},
      {270,	275,	280},
      {275,	280,	285},
      {280,	285,	290},
      {285,	290,	295},
      {290,	295,	300},
   }

   for i := 0; i < len(aArr); i++ {
      if i == len(aArr)/2 {
         fmt.Fprint(writer, "\n")
      }
      if t >= aArr[i][2] {
         fmt.Fprintf(writer, "%v ", 100)
         continue
      }
      if t >= aArr[i][1] {
         fmt.Fprintf(writer, "%v ", 300)
         continue
      }
      if t >= aArr[i][0] {
         fmt.Fprintf(writer, "%v ", 500)
         continue
      }
      fmt.Fprintf(writer, "%v ", 0)
   }
}
