package main

import (
  "fmt"
)

func insertionSort(a []int) {

  i, k := 0, 0

  for j := 1; j < len(a); j++ {

    k = a[j]
    i = j - 1

    for i > 0 && a[i] > k {
      a[i+1] = a[i]
      i = i - 1
    }
    a[i+1] = k
  }
}

func main() {

  nums := []int{1, 8, 3, 2, 1, 6, 4, 3, 2, 4, 6, 8}
  insertionSort(nums)

  fmt.Println(nums)
}
