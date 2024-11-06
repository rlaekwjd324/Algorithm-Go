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
   
   answer := ""
   for i := 0; i < n; i++ {
      var str string
      fmt.Fscanln(reader, &str)
      arr := strings.Split(str, "")
      number := ""
      count := 0
      for i := 0; i < len(arr); i++ {
         if arr[i] == "-" {
            continue
         }
         if count >= 10 {
            break
         }
         if count == 3 || count == 6 {
            number += "-"
         }
         if strings.Contains("1234567890", arr[i]) {
            number += arr[i]
         } else if strings.Contains("ABC", arr[i]) {
            number += "2"
         } else if strings.Contains("DEF", arr[i]) {
            number += "3"
         } else if strings.Contains("GHI", arr[i]) {
            number += "4"
         } else if strings.Contains("JKL", arr[i]) {
            number += "5"
         } else if strings.Contains("MNO", arr[i]) {
            number += "6"
         } else if strings.Contains("PQRS", arr[i]) {
            number += "7"
         } else if strings.Contains("TUV", arr[i]) {
            number += "8"
         } else if strings.Contains("WXYZ", arr[i]) {
            number += "9"
         }
         count++
      }
      answer += fmt.Sprintf("%s\n", number)
   }
   
   fmt.Fprintln(writer, answer)
}
