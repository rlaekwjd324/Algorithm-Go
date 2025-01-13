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

   new := ""
   for {
      var w string
      _, err := fmt.Fscan(reader, &w)
      if err != nil {
			break
		}

      for _, v := range w {
         n := int(v)
         if n >= 97 && n <= 122 {
            n -= 96
         } else if n >= 65 && n <= 91 {
            n -= 38
         } else if strings.Contains("1234567890", string(v)) {
            new += ("#"+string(v))
            continue
         } else {
            new += string(v)
            continue
         }
         m := fmt.Sprintf("%v", n)
         if n < 10 {
            m = "0"+m
         }
         new += m
      }
      new += " "
   }   
   fmt.Fprintln(writer, new)
}
