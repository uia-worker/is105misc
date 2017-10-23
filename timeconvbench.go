package main

import "fmt"
import "testing"

func BenchmarkWithShort(b *testing.B) {
    time12 := "12:00:00PM"
    hh := 12
    time24_hh := 0
    for i:=0;i < 100000000;i++ {
    if hh >= 1 && hh <= 12 {
      if time12[8:] == "PM" {
        time24_hh = hh
        if hh < 12 {
          time24_hh = hh + 12
        }
      }
    }
    hh = 12
    time24_hh = 0
    }
    _ = time24_hh
}

func BenchmarkWithLong(b *testing.B) {
    time12 := "12:00:00PM"
    hh := 12
    var time24_hh int
    for i:=0;i < 100000000;i++ {
    if hh >= 1 && hh <= 12 {
        if time12[8:] == "PM" {
            if hh < 12 {
                time24_hh = hh + 12
            } else {
                time24_hh = hh
            }
        } else if hh != 12 {
          time24_hh = hh
        }
    }
    hh = 12
    time24_hh = 0
    }
    _ = time24_hh
}

func main() {



    bwm := testing.Benchmark(BenchmarkWithShort)
    fmt.Println(bwm) // gives 176 ns/op

    bwl := testing.Benchmark(BenchmarkWithLong)
    fmt.Println(bwl) // gives 259 ns/op
}
