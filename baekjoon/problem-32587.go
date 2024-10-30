package main

import (
   "bufio"
	"fmt"
	"os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var t int
   fmt.Fscanln(reader, &t)
   
   var you string
   fmt.Fscanln(reader, &you)
   var guile string
   fmt.Fscanln(reader, &guile)
   
   arr1 := strings.Split(you, "")
   arr2 := strings.Split(guile, "")
   win := 0
   for j := 0; j < t; j++ {
      if arr1[j] == arr2[j] {
         continue
      }
      if arr1[j] == "R" && arr2[j] == "S" {
         win++
      } else if arr1[j] == "S" && arr2[j] == "P" {
         win++
      } else if arr1[j] == "P" && arr2[j] == "R" {
         win++
      } else {
         win--
      }
   }

   if win > 0 {
      fmt.Fprintln(writer, "victory")
      return
   }
   fmt.Fprintln(writer, "defeat")
}
