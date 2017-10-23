package main
import "fmt"

// test case
// Input: 140638725 436257910 953274816 734065819 362748590
// expected output: 1673711044 2486347135

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT

    size := 5
    numbers = make([]int64, size)
    for i := range numbers {
        fmt.Scanf("%d", &numbers[i])
        if (numbers[i] < 1) && (numbers[i] > 10e+9) {
            break
        }
    }
    fmt.Printf("len(numbers)=%d, cap(numbers)=%d\n", len(numbers), cap(numbers))
    var sum, min, max int64
    for j := range numbers {
        fmt.Println(numbers)
        for k := 0; k < len(numbers); k++ {
            if k != j {
                sum = sum + numbers[k]
            }
        }
        fmt.Println(sum)
        if min == 0 && max == 0 {
            min, max = sum, sum
        } else if min > sum {
            min = sum
        } else if max < sum {
            max = sum
        }
        fmt.Printf("min=%d\n", min)
        fmt.Printf("max=%d\n", max)
        sum = 0
    }

}
