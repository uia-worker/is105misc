package main

import (
  "fmt"
)

func main() {

  fmt.Println("assembler")

  a := [3]int{1, 2, 3}
  b := [3]int{4, 5, 6}
  n := len(a)

  sum := 0
  for i := 0; i <= n-1; i++ {
    sum = sum + a[i]*b[i]
  }

  fmt.Println(sum)

}

// R1 er akkumulator
// R2 er en base register
// ----------------------------------------------------
//   | n | sum | i | a0 | b0 | tmp | array a | array b |
// ----------------------------------------------------
//  ^                                ^         ^
//  R2                               a0        b0
//
// CLR R1
// STR R1,R2,2 sum=0
// STR R1,R2,3 i=0
// 
// LOAD R1,R2,4 R1=a0
// ADD  R1,R3,3 R1=a0+i
