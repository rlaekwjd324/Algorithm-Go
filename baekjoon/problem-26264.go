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

   var n int
   fmt.Fscanln(reader, &n)
   var word string
   fmt.Fscanln(reader, &word)
   
   words := strings.Split(word, "")

   count := 0
   index := 0
   for i := 0; i < n; i++ {
      if words[index] == "s" {
         count++
         index += 8
      }else if words[index] == "b" {
         count--
         index += 7
      }
   }
   if count == 0 {
      fmt.Fprintln(writer, "bigdata? security!")
   }else if count > 0 {
      fmt.Fprintln(writer, "security!")
   }else {
      fmt.Fprintln(writer, "bigdata?")
   }
}
