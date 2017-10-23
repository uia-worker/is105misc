package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n, tall int64
    fmt.Scanf("%v\n", &n)
    if n >= 1 && n <= 10e+5 {
        height := make([]int64, n)
        tall = 0
        counter := 0
        for i := range height {
            fmt.Scanf("%d", &height[i])
            if height[i] >= 1 && height[i] <= 10e+7 {
                // Calculation or comparision
                //for j := range height {
                    if height[i] > tall {
                        tall = height[i]
                        counter = 1
                        //fmt.Printf("isLarger %d, height[%d]=%d, tall=%d\n", counter, i, height[i], tall)
                    } else if height[i] == tall {
                        counter++
                        //fmt.Printf("isEqual %d, height[%d]=%d, tall=%d\n", counter, i, height[i], tall)
                    }
                //}
            }
        }
        fmt.Println(counter)
    }
}
