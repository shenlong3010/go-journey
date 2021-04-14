package main

import (
  "fmt"
  "sort"
)

func main() {
  fmt.Println(MinimumWaitingTime([]int{3, 2, 1, 2, 6}))
}


func MinimumWaitingTime(queries []int) int {
  sort.Ints(queries)

  totalWaitingTime := 0
  for idx, duration := range queries {
    queriesLeft := len(queries) - (idx+1)
    totalWaitingTime += duration * queriesLeft
  }
  return totalWaitingTime
}

