package main

import (
   "bufio"
   "fmt"
   "os"
   "sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

type TrieNode struct {
   children map[rune]*TrieNode
   isEnd    bool
}

func NewTrieNode() *TrieNode {
   return &TrieNode{children: make(map[rune]*TrieNode)}
}

type Trie struct {
   root *TrieNode
}

func NewTrie() *Trie {
   return &Trie{root: NewTrieNode()}
}

func main() {
   defer writer.Flush()

   var t int
   fmt.Fscan(reader, &t)

   for i := 0; i < t; i++ {
      var n int
      fmt.Fscan(reader, &n)

      numbers := make([]string, n)
      for j := 0; j < n; j++ {
         fmt.Fscan(reader, &numbers[j])
      }
      sort.Strings(numbers)

      trie := NewTrie()
      consistent := true

      for _, number := range numbers {
         if !trie.insert(number) {
            consistent = false
            break
         }
      }
      if !consistent {
         fmt.Fprintf(writer, "%v\n", "NO")
         continue
      }
      fmt.Fprintf(writer, "%v\n", "YES")
   }
}

func (t *Trie) insert(number string) bool {
   node := t.root
   for _, digit := range number {
      if node.isEnd {
         return false
      }
      if _, exists := node.children[digit]; !exists {
         node.children[digit] = NewTrieNode()
      }
      node = node.children[digit]
   }
   if len(node.children) > 0 {
      return false
   }
   node.isEnd = true
   return true
}

