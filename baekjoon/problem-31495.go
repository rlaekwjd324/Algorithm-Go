package main

import (
   "bufio"
   "fmt"
   "os"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   scanner := bufio.NewScanner(os.Stdin)

	// 한 줄 입력받기
	scanner.Scan()
	s := scanner.Text()

   if len(s) <= 2 {
     fmt.Fprintln(writer, "CE")
     return
   }
   for i, _ := range s {
     if i == 0 || i == len(s) - 1 {
         if string(s[i]) != "\"" {
             fmt.Fprintln(writer, "CE")
             return
         }
     }
   }
   fmt.Fprintln(writer, s[1:len(s)-1])
}
