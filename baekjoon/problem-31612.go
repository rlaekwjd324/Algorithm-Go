package main

import (
   "fmt"
   "strings"
)

func main() {
   var t int
   fmt.Scanf("%v", &t)
   var word string
   fmt.Scanf("%v", &word)
   word = strings.TrimSpace(word)
   wordArr := strings.Split(word, "")
   
   var sum = 0
   for i := 0; i < t; i++ {
      if wordArr[i] == "o"{
         sum += 1
      }else {
         sum += 2
      }
   }
   fmt.Println(sum)
}
