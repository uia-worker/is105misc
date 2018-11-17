package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var hashtag string
    var size int
    fmt.Scanf("%v\n", &size)
    for i := size; i > 0; i-- {
        for j := 0; j < i-1; j++ {
            hashtag = hashtag + fmt.Sprintf("%s", " ")
        }
        //fmt.Println(len(hashtag))
        //if len(hashtag) < size {

        fmt.Printf("size-len(hashtag)=%d\n", size-len(hashtag))
        h := 0
        for h < (size-len(hashtag)) {
            fmt.Printf("h=%d\n", h)
            hashtag = hashtag + fmt.Sprintf("%s", "#")
        }
        //}
        //fmt.Printf("%s\n", hashtag)
        hashtag = ""
    }
}
