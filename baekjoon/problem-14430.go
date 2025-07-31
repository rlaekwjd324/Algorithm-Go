package main

import (
   "bufio"
   "fmt"
   "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var (
   n, m int // 미로의 크기
   arr [][]int // 입력받은 미로
   dp [][]int // [i][j]에서의 최댓값 저장
)

func main() {
   defer writer.Flush()

   fmt.Fscan(reader, &n, &m)

   // 입력값
   arr = make([][]int, n)
   dp = make([][]int, n)
   for i := 0; i < n; i++ {
      arr[i] = make([]int, m)
      dp[i] = make([]int, m)
      for j := 0; j < m; j++ {
         fmt.Fscan(reader, &arr[i][j])
      }
   }

   // 출력값
   fmt.Fprintf(writer, "%v\n", visitMap())
}

func visitMap() int {
   for i := 0; i < n; i++ {
      for j := 0; j < m; j++ {
         if i > 0 {
            dp[i][j] = max(dp[i][j], dp[i-1][j])
         }
         if j > 0 {
            dp[i][j] = max(dp[i][j], dp[i][j-1])
         }
         dp[i][j] += arr[i][j]
      }
   }
   
   return dp[n-1][m-1]
}

func getBiggerNumber(a, b int) int {
   if a > b {
      return a
   }
   return b
}
