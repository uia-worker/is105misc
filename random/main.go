package main

import ( 
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("BEGIN")
	list := rand.Perm(20)
		//for i, _ := range list {
    	//list[i]++
	//}
	res := fmt.Sprintf("%v", list)
	fmt.Println(res)
}