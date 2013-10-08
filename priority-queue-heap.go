// Priority Queue Implementation using a min-heap with a heapsort routine
// to demonstrate priority queue.
// 
// Correctness not guaranteed. Use at your own risk.
// 
// Eddie Blundell eblundell@gmail.com
// (Mostly translated from The Algorithm Design Manual - Steven S, Skiena)

package main

import "fmt"
import "math/rand" // for testing

const QueueSize = 100000
type QueueItem int

type priority_queue struct {
  n int
  q [QueueSize]QueueItem
}

func (q *priority_queue) Insert(x QueueItem) {

  if q.n >= QueueSize {
    // Queue Overflow
    return
  }

  q.n ++
  q.q[q.n] = x
  q.BubbleUp(q.n)
}

func (q *priority_queue) Swap(a int, b int) {
  
  tmp := q.q[a]
  q.q[a] = q.q[b]
  q.q[b] = tmp
}

func (q *priority_queue) BubbleUp(p int) {

  if q.Parent(p) == -1 {
    return
  }

  if q.q[q.Parent(p)] > q.q[p] {
    q.Swap(p, q.Parent(p))
    q.BubbleUp(q.Parent(p))
  }
}

func (q *priority_queue) YoungChild(n int) int {
  if n == 0 {
    return 1
  }
  return 2 * n
}

func (q *priority_queue) Parent(n int) int {

  if n == 0 {
    return -1
  }

  return int(n/2)
}

func (q *priority_queue) ExtractMin() QueueItem {

  var min QueueItem = -1

  if q.n <= 0 {
    // Queue Empty
  } else {

    min = q.q[0]
    q.q[0] = q.q[q.n]
    q.q[q.n] = 0
    q.n--

    q.BubbleDown(0)
  }

  return min
}

func (q *priority_queue) BubbleDown(p int) {
  
  c := q.YoungChild(p)
  min_index := p

  for i := 0; i <= 1; i++ {
    if (c+i) <= q.n {
      if q.q[min_index] > q.q[c+i] {
        min_index = c + i
      }
    }
  }

  if min_index != p {
    q.Swap(p, min_index)
    q.BubbleDown(min_index)
  }
}

func MakeHeap(items []QueueItem) *priority_queue {

  var p_queue priority_queue
  p_queue.n = -1

  for _,item := range items {
    p_queue.Insert(item)
  }

  return &p_queue
}

func Heapsort(items []QueueItem) {

  p_queue := MakeHeap(items)

  tmp := p_queue.ExtractMin()

  for tmp >= 0 {
    fmt.Println(tmp)
    tmp = p_queue.ExtractMin()
  } 
}

func main() {

  items := make([]QueueItem, QueueSize-1)

  for k,_ := range items {
    items[k] = QueueItem(rand.Intn(100))
  }

  Heapsort(items)

}
