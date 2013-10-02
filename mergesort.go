package main

import "fmt"

// Merge sort using channels

func Merge(s []int, p,r,q int) {

  ch1, ch2 := make(chan int), make(chan int)

  for i := p; i <= r; i++ {
    ch1 <- s[i]
  }

  for i := r+1; i <= q; i++ {
    ch2 <- s[i]
  }
  
  i := p

  for (len(ch1) > 0) || (len(ch2) > 0) {
    
    if <-ch1 <= <-ch2 {
      s[i] = <-ch1
    } else {
      s[i] = <-ch2
    }

    i = i + 1
  }

}

func MergeSort(s [] int, p,q int) {

  var r int

  if p < q {
    
    r = (p+q)/2
    MergeSort(s, p, r)
    MergeSort(s, r + 1, q)
    Merge(s, p, r, q)
  }
}

func main() {

  s := []int{1,4,1,3,2,9,4,2}
  MergeSort(s,0, len(s))

  fmt.Println(s)
}


