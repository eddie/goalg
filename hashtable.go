// Quick hack to practice Golang
//
// A hash table implementation using chaining to
// handle hash collisions
//
// Eddie Blundell eblundell@gmail.com

package main

import "fmt"
import "math"
import "crypto/rand" // For tests

type ListItem struct {
  Key string
  Value int
}

type List struct {
  Prev *List
  Item *ListItem
  Next *List
}

type HashTable struct {
  Items [127]*List
}

func (l *List) Insert(x *ListItem) {
  t := l.Next
  l.Next = &List{l, x, t}
}

func (l *List) Search(key string) *List {
  if l == nil {
    return nil
  }

  if l.Item.Key == key {
    return l
  } 
  
  return l.Next.Search(key)
}

func (l *List) len() int {

  tmp := l
  count := 0

  for tmp != nil {
    count ++
    tmp = tmp.Next
  }
  return count
}

func (l *List) Delete(key string) {
  
  node := l.Search(key)

  node.Prev.Next = node.Next
  node.Next.Prev = node.Prev
}

// O(len(str)) hashing function
func hash(str string) (hash int64) {

  hash = 0
  m := len(str)

  for i, c := range str {
    hash += int64(math.Pow(26, float64(m-(i+1)) ) * float64(c))
  }

  return hash
}

func (t *HashTable) Set(key string, value int) {

  hash := hash(key) % 127

  if t.Items[hash] == nil {
    t.Items[hash] = &List{nil, &ListItem{key,value}, nil}
  } else {
    t.Items[hash].Insert(&ListItem{key, value})
  }
}

func (t *HashTable) Get(key string) int {
  
  hash := hash(key) % 127

  if t.Items[hash] == nil {
    return -1
  }

  return t.Items[hash].Search(key).Item.Value
}

func (t *HashTable) Stats() {
  
  fmt.Println("=================================")
  total := 0

  for i, x := range t.Items {

    if x != nil {
      fmt.Printf("Bucket %d %d \n", i, x.len())
      total += x.len()
    }
  }
  fmt.Println("=================================")
  fmt.Printf("Items: %d\n", total)
  fmt.Println("=================================")
}

func randString(n int) string {
  const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
  var bytes = make([]byte, n)
  rand.Read(bytes)
  for i, b := range bytes {
    bytes[i] = alphanum[b % byte(len(alphanum))]
  }
  return string(bytes)
}

func main() {

  // Our hash table
  var ht HashTable

  for i := 0; i < 500000; i ++ {
    
    str := randString(5)
    ht.Set(str,i)
    fmt.Printf("String: %s Value: %d\n", str, ht.Get(str))
  }

  ht.Stats()
}
