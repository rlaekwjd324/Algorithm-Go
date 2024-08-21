package main

import (
   "fmt"
   "bufio"
   "strings"
   "os"
   "strconv"
)

func main() {
	 reader := bufio.NewReader(os.Stdin)
   input, _ := reader.ReadString('\n')
   input = strings.TrimSpace(input)
   nums := strings.Split(input, " ")
   n, _ := strconv.Atoi(nums[0])
   k, _ := strconv.Atoi(nums[1])
   a, _ := strconv.Atoi(nums[2])
   b, _ := strconv.Atoi(nums[3])
   walk := (n-1) * a
   ele := (n-1+k-1)*b
   fmt.Printf("%v %v\n", ele, walk)
}
