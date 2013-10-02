package main

import (
  "fmt"
)

type Tree struct {

  Left *Tree
  Value int
  Right *Tree
}

func (t *Tree) Insert(v int) *Tree {
  if t == nil {
    return &Tree{nil, v, nil}
  }

  if v < t.Value {
    t.Left = t.Left.Insert(v)
    return t.Left
  }

  t.Right = t.Right.Insert(v)
  return t.Right
}

// Depth first in-order traversal
func (t *Tree) Traverse(ch chan int) {

  if t == nil {
    return
  }

  t.Left.Traverse(ch)
  ch <- t.Value
  t.Right.Traverse(ch)
}

func (t *Tree) Search(x int) *Tree {

  if t == nil {
    return nil 
  }

  if t.Value == x {
    return t
  }

  if x < t.Value {
    return t.Left.Search(x)
  }
  
  return t.Right.Search(x)
}

// Post-order traversal to calculate tree depth
func (t *Tree) Depth() int {
  
  if t == nil {
    return 0
  }

  var left,right int
  left = t.Left.Depth()
  right = t.Right.Depth()

  if left > right {
    return left + 1
  }

  return right + 1
}

func (t *Tree) Min() *Tree {
  
  var min *Tree

  if t == nil {
    return nil
  }

  min = t
  for min.Left != nil {
    min = min.Left
  }

  return min
}

// Pre-order traversal for graphviz
func (t *Tree) GraphVizAux() {

  if(t == nil) {
    return
  }

  if(t.Left != nil) {
    fmt.Printf("  %d->%d; \n", t.Value, t.Left.Value)
  }
  if(t.Right != nil) {
    fmt.Printf("  %d->%d; \n", t.Value, t.Right.Value)
  }

  t.Left.GraphVizAux();
  t.Right.GraphVizAux();
}

func (t *Tree) GraphViz() {

  fmt.Println("digraph BST {")
  fmt.Println("graph [ordering=\"out\"];")
  t.GraphVizAux() 
  fmt.Println("labelloc=\"t\";")
  fmt.Printf("label=\"Binary Search Tree with Depth %d\";\n", t.Depth())
  fmt.Println("}")
}

func main() {

  var t,tmp,tmp2 *Tree

  t = t.Insert(2)
  t.Insert(1)
  tmp = t.Insert(7)
  tmp2 = tmp.Insert(4)
  tmp = tmp.Insert(8)
  tmp2.Insert(3)
  tmp = tmp2.Insert(6)
  tmp.Insert(5)

  /*
  ch := make(chan int)
  go func() {
    t.Traverse(ch)
    close(ch)
  }()
  for i := range ch {
    fmt.Println(i)
  }
  */
  t.GraphViz()

}
