package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var size int
    var p, n, z int
    var numbers []int
    p,n,z = 0, 0, 0
    fmt.Scanf("%v\n", &size)
    numbers = make([]int, size)
    for i := range numbers {
      fmt.Scanf("%d", &numbers[i])
      if numbers[i] > 0 {
        p++
      } else if numbers[i] < 0 {
          n++
      } else {
          z++
      }
    }
    fmt.Printf("%f\n",float32(p)/float32(size))
    fmt.Printf("%f\n",float32(n)/float32(size))
    fmt.Printf("%f\n",float32(z)/float32(size))

}
