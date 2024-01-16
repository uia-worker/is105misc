package main
import (
   "fmt"
   "math"
)
func printPrimeNumbers(num1, num2 int) {
   if num1 < 2 || num2 < 2 {
      fmt.Println("Numbers must be greater than 2.")
      return
      
   }
   for num1 <= num2 {
      isPrime := true
      for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
         if num1%i == 0 {
            isPrime = false
            break
            
         }
      }
      if isPrime {
         fmt.Printf("%d ", num1)
      }
      num1++
      
   }
   
}
func main() {
   var range1 int = 4
   var range2 int = 100
   fmt.Println("The range from which prime numbers have to be searched is from", range1, "to", range2)
   fmt.Println("The following prime numbers are: ")
   printPrimeNumbers(4, 100)
}

