package main

import (
  "fmt"
)

func main() {
  var s []byte
  s = make([]byte, 5)
  s[0] = 'J'
	s[1] = 'a'
	s[2] = 'n'
	s[3] = 'i'
	s[4] = 's'
  t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
  nb := copy(t, s) // t - destination, s -source
  fmt.Printf("number of bytes copied: %d\n", nb)
  debug(s)

  s = t // get the pointer s point to the new slice
  s = s[:len(s)+1] // need to grow slice first to its capacity (or less)
  s[5] = ' ' // panic if you not grow the slice first
  debug(s)
  s = appendByte(s, 'G', 'a', 'i', 'l')
  debug(s)
  s = appendByte(s, 'i', 's', '.')
  debug(s)
}

func appendByte(s []byte, data ...byte) []byte {
  m := len(s)
  n := m + len(data)
  if n > cap(s) { // if necessary, reallocate
    // allocate double what's needed, for future growth
    newSlice := make([]byte, (n+1)*2) // allocate the memory
    copy(newSlice, s) // put the data into the new allocation
    s = newSlice // direct the pointer to the new slice
  }
  s = s[0:n] // extend capacity to what is necessary?
  copy(s[m:n], data)
  return s
}

func copy(t []byte, s []byte) int {
  count := 0
  for i := range s {
    t[i] = s[i]
    count++
  }
  //fmt.Printf("copy: len(t) = %d\n", len(t))
  //fmt.Printf("copy: cap(t) = %d\n", cap(t))
  return count
}

func debug (s []byte) {
  fmt.Printf("len(s) = %d\n", len(s))
  fmt.Printf("cap(s) = %d\n", cap(s))
  fmt.Printf("%c\n", s)
}
