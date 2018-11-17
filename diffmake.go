package main

import "fmt"
import "testing"

func BenchmarkWithMake(b *testing.B) {
    m0 := make(map[int]int, b.N)
    for i := 0; i < b.N; i++ {
        m0[i] = 1000
    }
}

func BenchmarkWithLitteral(b *testing.B) {
    m1 := map[int]int{}
    for i := 0; i < b.N; i++ {
        m1[i] = 1000
    }
}

func main() {
    bwm := testing.Benchmark(BenchmarkWithMake)
    fmt.Println(bwm) // gives 176 ns/op

    bwl := testing.Benchmark(BenchmarkWithLitteral)
    fmt.Println(bwl) // gives 259 ns/op
}
