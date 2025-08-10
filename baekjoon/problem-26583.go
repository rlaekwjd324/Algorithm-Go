package main

import (
	"fmt"
	"bufio"
	"os"
   "strings"
   "strconv"
)

// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
// var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	// defer writer.Flush()
   scanner := bufio.NewScanner(os.Stdin)
   for scanner.Scan() {
      line := strings.TrimSpace(scanner.Text())
      if line == "" {
         return
      }
      numstr := strings.Split(line, " ")
      nums := []int{1}
      for j := 0; j < len(numstr); j++ {
         n, _ := strconv.Atoi(numstr[j])
         nums = append(nums, n)
      }
      nums = append(nums, 1)
      arr := make([]string, len(numstr))
      for j := 0; j < len(numstr); j++ {
         arr[j] = fmt.Sprintf("%v", nums[j]*nums[j+1]*nums[j+2])
      }
      fmt.Printf("%v\n", strings.Join(arr, " "))
   }
}
