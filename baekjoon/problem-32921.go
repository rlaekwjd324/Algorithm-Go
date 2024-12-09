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
   fmt.Fscanln(reader, &n)

   preH2 := 0
   preM2 := 0
   preH1 := 0
   preM1 := 0
   min := 2
   hs := make([]int, n)
   ms := make([]int, n)
   for i := 0; i < n; i++ {
      var time string
      fmt.Fscanln(reader, &time)
      h, _ := strconv.Atoi(strings.Split(time, ":")[0])
      m, _ := strconv.Atoi(strings.Split(time, ":")[1])
      hs[i] = h
      ms[i] = m
   }
   for i := 0; i < n; i++ {
      if i >= 2 {
         tH := hs[i] - preH2
         tM := ms[i] - preM2
         if tH > 0 {
            tM += tH * 60
         }
         if tM <= 10 {
            min = 0
            break
         }
      }
      if i >= 1 {
         tH := hs[i] - preH1
         tM := ms[i] - preM1
         if tH > 0 {
            tM += tH * 60
         }
         if tM <= 10 {
            min = 1
         }
         preH2 = preH1
         preM2 = preM1
      }
      preH1 = hs[i]
      preM1 = ms[i]
   }

   fmt.Fprintln(writer, min)
}
