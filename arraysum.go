package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var size int
    var array []int
    // Input is (with the line shift)
    // 6
    // 1 2 3 4 10 11
    // Output should be a sum of all 6 elements
    fmt.Scanf("%v\n", &size)
    array = make([]int, size)
    sum := 0
    for i := range array {
        fmt.Scanf("%d", &array[i])
        sum = sum + array[i]
    }
    fmt.Printf("%d\n", sum)
}
