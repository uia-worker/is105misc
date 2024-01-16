package main

import "fmt"

func main() {
  hello := "Hello"
  for {
	  hello = hello + "Hello"
	  fmt.Println(hello)
  }
}

// go tool asm [flags] file