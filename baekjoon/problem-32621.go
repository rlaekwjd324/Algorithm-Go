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

   var n string
   fmt.Fscanln(reader, &n)
   if !strings.Contains(n, "+") {
      fmt.Fprintln(writer, "No Money")
      return
   }
   
   arr := strings.Split(n, "+")
   if len(arr) > 2 {
      fmt.Fprintln(writer, "No Money")
      return
   }

   a, err := strconv.Atoi(arr[0])
   if err != nil {
      fmt.Fprintln(writer, "No Money")
      return
   }
   b, err := strconv.Atoi(arr[1])
   if err != nil {
      fmt.Fprintln(writer, "No Money")
      return
   }
   if a <= 0 || b <= 0 || a != b {
      fmt.Fprintln(writer, "No Money")
      return
   }
   if strconv.Itoa(a) != arr[0] || strconv.Itoa(b) != arr[1] {
      fmt.Fprintln(writer, "No Money")
      return
   }
   fmt.Fprintln(writer, "CUTE")
}
