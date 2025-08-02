package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var(
   n, m int
   arr [][]string
   min = 1000000
)

type Node struct {
   x, y int
   breakWall int
}

func main() {
   defer writer.Flush()

   fmt.Fscan(reader, &n, &m)
   
   arr = make([][]string, n)
   
   for i := 0; i < n; i++ {
      var w string
      fmt.Fscan(reader, &w)
      arr[i] = strings.Split(w, "")
   }

   fmt.Fprintf(writer, "%v\n", bfs())
}

func bfs() int {
   dx := []int{-1, 1, 0, 0}
   dy := []int{0, 0, -1, 1}

   // 거리 저장용 배열 (벽 부술때/안부술때)
   visited := make([][][]int, n)
   for i := 0; i < n; i++ {
		visited[i] = make([][]int, m)
      	for j := 0; j < m; j++ {
         visited[i][j] = []int{0, 0}
      }
	}

   queue := []Node{{0, 0, 0}}
   visited[0][0][0] = 1

   for len(queue) > 0 {
      current := queue[0]
      queue = queue[1:]

      for i := 0; i < 4; i++ {
         nx := current.x + dx[i]
         ny := current.y + dy[i]

         if nx >= 0 && nx < n && ny >= 0 && ny < m {
            // 벽 아니면
            if arr[nx][ny] == "0" && visited[nx][ny][current.breakWall] == 0 {
               visited[nx][ny][current.breakWall] = visited[current.x][current.y][current.breakWall] + 1
               queue = append(queue, Node{nx, ny, current.breakWall})
            }
            // 벽이고 아직 안 부쉈으면
            if arr[nx][ny] == "1" && current.breakWall == 0 && visited[nx][ny][1] == 0 {
               visited[nx][ny][1] = visited[current.x][current.y][current.breakWall] + 1
               queue = append(queue, Node{nx, ny, 1})
            }
         }
      }
   }

   a0 := visited[n-1][m-1][0]
   a1 := visited[n-1][m-1][1]
   // 도달 불가능
   if a0 == 0 && a1 == 0 {
      return -1
   }
   if a0 == 0 {
      return a1
   }
   if a1 == 0 {
      return a0
   }
   if a0 > a1 {
      return a1
   }
   return a0
}
