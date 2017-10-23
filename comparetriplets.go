package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var alice, bob []int
    palice, pbob := 0,0
    alice = make([]int, 3)
    bob = make([]int, 3)
    
    for i := range alice {
        fmt.Scanf("%d", &alice[i])
    }
    for i := range bob {
        fmt.Scanf("%d", &bob[i])
    }
    //fmt.Printf("%v", alice)
    //fmt.Printf("%v", bob)

    // [clarity, originality, difficulty]
    // between 1 and 100
    for p := 0; p < 3; p++ {
        if (alice[p] >= 1) && (alice[p] <= 100) && (bob[p] >= 1) && (bob[p] <= 100) {
            if alice[p] > bob[p] {
                palice++
            } else if bob[p] > alice[p] {
                pbob++
            }
        }
    }

    fmt.Println(palice,pbob)

}
