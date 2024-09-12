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

   var arr = []string{}
   arr = append(arr, "A B C D E F G H J L M")
   arr = append(arr, "A C E F G H I L M")
   arr = append(arr, "A C E F G H I L M")
   arr = append(arr, "A B C E F G H L M")
   arr = append(arr, "A C E F G H L M")
   arr = append(arr, "A C E F G H L M")
   arr = append(arr, "A C E F G H L M")
   arr = append(arr, "A C E F G H L M")
   arr = append(arr, "A C E F G H L M")
   arr = append(arr, "A B C F G H L M")
   
   count := len(arr[n-1])/2+1
   fmt.Fprintln(writer, count)
   fmt.Fprintln(writer, arr[n-1])
}
