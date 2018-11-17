package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var time12 string
    var hh,mm,ss,time24_hh int
    fmt.Scanf("%v\n", &time12)
    fmt.Sscanf(time12, "%d:%d:%d", &hh, &mm, &ss)
    // Her må legge inn sjekk for mm og ss [0,59]
    if hh >= 1 && hh <= 12 {
      /*
        if time12[8:] == "PM" {
            if hh < 12 {
                time24_hh = hh + 12
            } else {
                time24_hh = hh
            }
        } else if hh != 12 {
          time24_hh = hh
        }
      */

      if time12[8:] == "PM" {
        time24_hh = hh
        if hh < 12 {
          time24_hh = hh + 12
        }
      }
        //fmt.Println(hh, mm, ss)
        // Bør være sjekk at time24_hh [0,23]
        fmt.Printf("%02d:%02d:%02d\n", time24_hh, mm, ss)
    }
}
