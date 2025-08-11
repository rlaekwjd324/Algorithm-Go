package main

import (
	"fmt"
	"bufio"
	"os"
  "strings"
)

func main() {
   scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	w := scanner.Text()
   
   c := 0 
   for scanner.Scan() {
      line := scanner.Text()
      c += strings.Count(line, w)
   }
   fmt.Println(c)
}
