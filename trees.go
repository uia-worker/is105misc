package main

import (
  "fmt"
  "math/rand"
)

// Here the tree is a binary tree with integer values
type Tree struct {
  Left *Tree  // explain the pointer here
  Value int
  Right *Tree
}

// Creates a new, random binary tree holding the values 1k, 2k, ..., nk
func New(n,k int) *Tree { // returns a pointer to the struct Tree
  var t *Tree
  for _, v := range rand.Perm(n) { // Check the for _, v explanation
    t = insert(t, (1+v)*k)
  }
  return t
}

// Recursion is used here
// 3 possible situations are handled here
// (1) "zero" situation, if the pointer t of type Tree points to nothing (nil)
// (2) if the value in the tree struct is larger than the value to insert into
//     the tree (a formation in memory, build up by pointer references)
// (3) if none of (1) and (2) is true, the third situation yields, - the value
//     is inserted as a right node of the struct pointed to by t of type *Tree
func insert(t *Tree, v int) *Tree {
  if t == nil {
    return &Tree{nil, v, nil} // break the subroutine if tree is not made
  }
  if v < t.Value {
    t.Left = insert(t.Left, v)
    return t // break the subroutine if the the value is less than tree value
  }
  t.Right = insert(t.Right, v)
  return t
}

func Walk(t *Tree, ch chan int) {
  if t == nil {
    return
  }
  Walk(t.Left, ch)
  ch <- t.Value
  Walk(t.Right, ch)
}

func Walker(t *Tree) <-chan int {
  ch := make(chan int)
  go func() {
    Walk(t, ch)
    close(ch)
  }()
  return ch
}

func main() {
  t1 := New(100, 1)
  fmt.Printf("%v\n", t1.Value)
  ch := Walker(t1)
  // Read from channel
  for {
    v1, ok1 := <-ch
    fmt.Println(v1, ok1)
  }
  fmt.Println("-------------")
  for _, v := range rand.Perm(7) {
    fmt.Printf("%v\n", v)
    fmt.Printf("%v\n", (1+v)*3)
  }
}
