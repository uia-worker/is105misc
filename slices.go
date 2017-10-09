package main

import (
	"fmt"
)

func main() {
	d := []byte{'r','o','a','d'}
	fmt.Printf("%p\n", d)
	e := d[2:]
	fmt.Printf("%p\n", e)
	e[1] = 'm'

	fmt.Printf("%c\n", d)
	fmt.Printf("%c\n", e)

	fmt.Printf("len(d) = %d\n", len(d))
	fmt.Printf("cap(d) = %d\n", cap(d))

	fmt.Printf("len(e) = %d\n", len(e))
	fmt.Printf("cap(e) = %d\n", cap(e))

	var s []byte
	s = make([]byte, 5)
	s[0] = 'J'
	s[1] = 'a'
	s[2] = 'n'
	s[3] = 'i'
	s[4] = 's'
	fmt.Printf("%c\n", s)
	fmt.Printf("len(s) = %d\n", len(s))
	fmt.Printf("cap(s) = %d\n", cap(s))
	s = s[2:4]
	fmt.Printf("%c\n", s)
	fmt.Printf("len(s) = %d\n", len(s))
	fmt.Printf("cap(s) = %d\n", cap(s))
	s = s[:cap(s)]
	fmt.Printf("len(s) = %d\n", len(s))
	fmt.Printf("cap(s) = %d\n", cap(s))
	s = s[:cap(s)+1] // panic: runtime error: slice bounds out of range
}

/*
0x10410020
0x10410022
[r o a m]
[a m]
len(d) = 4
cap(d) = 4
len(e) = 2
cap(e) = 2
[J a n i s]
len(s) = 5
cap(s) = 5
[n i]
len(s) = 2
cap(s) = 3
len(s) = 3
cap(s) = 3
*/
