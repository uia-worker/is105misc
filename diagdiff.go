package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var size int
    fmt.Scanf("%v\n", &size)
    matrix := [][]int{}
    row := make([]int, size)
    dsum, osum := 0, 0
    for _ = range row {
        for i := range row {
         fmt.Scanf("%d", &row[i])
            if (row[i] <= -100) && (row[i] >= 100) {
                break
            }
        }
        matrix = append(matrix, row)
        row = make([]int, size)
    }

    for j := 0; j < size; j++ {
        dsum = dsum + matrix[j][j]
        osum = osum + matrix[j][size-j-1]
    }
    if (dsum-osum) < 0 {
        fmt.Println(-(dsum-osum))
    } else {
        fmt.Println(dsum-osum)
    }
}
