package main

import (
   "fmt"
   "strings"
)

func main() {
   var n int
   fmt.Scanln(&n)
   var str string
   fmt.Scanln(&str)
   
   for {
      if strings.Contains(str, "PS4") {
         index := strings.Index(str, "PS4")
         str = str[:index+2]+str[index+3:]
      } else if strings.Contains(str, "PS5") {
         index := strings.Index(str, "PS5")
         str = str[:index+2]+str[index+3:]
      } else {
         break
      }
   }
   
   fmt.Println(str)
}
