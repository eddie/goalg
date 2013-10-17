package main

import (
  "fmt"
  "sort"
)

// Recursive Binary Search 
// Runs in O(lg n) worst case time

func binarySearch(a []int, p, r, v int) int {

  if p <= r {

    q := (p + r) / 2

    if a[q] == v {
      return q
    }

    if a[q] > v {
      return binarySearch(a, p, q-1, v)
    } else {
      return binarySearch(a, q+1, r, v)
    }

    fmt.Println(q)
  }

  return -1
}

func main() {

  nums := []int{1, 8, 3, 2, 1, 6, 4, 3, 2, 4, 6, 8}
  sort.Ints(nums)

  index := binarySearch(nums, 0, len(nums), 8)
  fmt.Println("Found 8 at index:", index)

  fmt.Println("Hello World")
}
