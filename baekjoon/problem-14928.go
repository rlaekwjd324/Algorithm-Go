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
   
   var s string
   fmt.Fscan(reader, &s)

   MOD := 20000303
   
   result := 0
   for i := 0; i < len(s); i++ {
      result = (result*10 + int(s[i]-'0')) % MOD
   }
   fmt.Fprintf(writer, "%v\n", result)
}
