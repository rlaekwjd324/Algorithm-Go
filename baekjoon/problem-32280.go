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

   var n int
   fmt.Fscan(reader, &n)

   students := []string{}
   teacher := ""

   for i := 0; i < n+1; i++ {
      var time, job string
      fmt.Fscan(reader, &time, &job)
      if job == "student" {
         students = append(students, time)
      }else {
         teacher = time
      }
   }
   var school string
   fmt.Fscan(reader, &school)

   count := 0
   for i := 0; i < n; i++ {
      if isLate(students[i], school) && isLate(students[i], teacher) {
         count++
      }
   }

   fmt.Fprintln(writer, count)
}

func isLate(first string, second string) bool {
   firstH, _ := strconv.Atoi(strings.Split(first, ":")[0])
   firstM, _ := strconv.Atoi(strings.Split(first, ":")[1])
   secondH, _ := strconv.Atoi(strings.Split(second, ":")[0])
   secondM, _ := strconv.Atoi(strings.Split(second, ":")[1])
   if firstH > secondH {
      return true
   }
   if firstH < secondH {
      return false
   }
   if firstM >= secondM {
      return true
   }
   return false
}
