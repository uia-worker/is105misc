/*
44
84
94
21
0
18
100
18
62
30
61
53
0
43
2
29
53
61
40
14
4
29
98
37
23
46
9
79
62
20
38
51
99
59
47
4
86
61
68
17
45
6
1
95
95

85
95
21
0
18
100
18
62
30
61
55
0
45
2
29
55
61
40
14
4
29
100
37
23
46
9
80
62
20
40
51
100
60
47
4
86
61
70
17
45
6
1
95
95
*/

package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var grade, number int
    var fgrades []int
    fmt.Scanf("%v\n", &number)
    if number >= 1 && number <= 60 {
        fgrades = make([]int, number)
        for i := 0; i < number; i++ {
            fmt.Scanf("%d\n", &grade)
            if grade >= 0 && grade <= 100 {
                //fmt.Println(grade % 5)
                if (grade % 5) >= 3 && grade >= 38 {
                    //fmt.Println(grade)
                    fgrades[i] = grade - (grade % 5) + 5
                } else {
                    fgrades[i] = grade
                }
                fmt.Println(fgrades[i])
            } // end if check for legal values of grades
        }
    } // end if check for legal values of number of grades

}
