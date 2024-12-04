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

   n, err := reader.ReadString('\n')
   if err != nil {
		return
	}

   if strings.Contains(n, "(") {
      arr := strings.Split(n, " (")
      fmt.Fprintln(writer, arr[0])
      m := strings.Split(arr[1], ")")[0]
      fmt.Fprintln(writer, m)
   } else {
      fmt.Fprint(writer, n)
      fmt.Fprintln(writer, "-")
   }
}
