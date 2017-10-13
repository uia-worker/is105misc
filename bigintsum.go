package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 // int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
 // int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
    var size int
    var sum int64
    var bignumbers []int64
    fmt.Scanf("%v\n", &size)
    if (size >= 1) && (size <= 10) {
        bignumbers = make([]int64, size)
        sum = 0
        for i := range bignumbers {
            fmt.Scanf("%d", &bignumbers[i])
            if (bignumbers[i] >= 0) && (bignumbers[i] <= 10e+10) {
                sum = sum + bignumbers[i]
            }
        }
    }
    fmt.Printf("%d\n", sum)
}
