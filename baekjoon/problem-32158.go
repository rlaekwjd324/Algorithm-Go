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
   fmt.Fscan(reader, &n)
   var word string
   fmt.Fscan(reader, &word)
   
   var arr = []string{}
   for i := 0; i < n; i++ { 
      arr = append(arr, "")
   }
   temp := strings.Split(word, "")
   for i := 0; i < n; i++ {
      if arr[i] != "" {
         continue
      }
      if temp[i] != "P" && temp[i] != "C" {
         arr[i] = temp[i]
         continue
      }
      if temp[i] == "P" {
         for j := i+1; j<n; j++ {
            if arr[j] == "" && temp[j] == "C" {
               arr[i] = "C"
               arr[j] = "P"
               break
            }
         }
      } else {
         for j := i+1; j<n; j++ {
            if arr[j] == "" && temp[j] == "P" {
               arr[i] = "P"
               arr[j] = "C"
               break
            }
         }
      }
   }
   newWord := ""
   for i := 0; i < n; i++ {
      newWord += arr[i]
      if arr[i] == "" {
         newWord += temp[i]
      }
   }

   fmt.Fprintln(writer, newWord)
}
