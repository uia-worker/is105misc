package main
import "fmt"

// n cities, [1,n]
// m bidiractional roads connect cities
// A citizen has access to a library if:
// (a) their city contains a library
// (b) they can travel by road from their city to a city
//     containing a library
// Cost of repairing any road is c[road] dollars
// Cost to build a library in a city is c[lib] dollars
// q - queries, each query consists of a map of
//     Hackerland and value c[lib] and c[road]
// For each query find the minimum cost of making
// libraries accessible to all citizens and print it
// on a new line

func main() {
  // INPUT: The first line contains single integer q,
  //        number of queries
  //        The second line contains four space-separated
  //        integers describing respective values of
  //        n - number of cities
  //        m - the number of roads
  //        c[lib] - the cost to build a library
  //        c[road] - the cost to repair a road
  var q, n, m, clib, croad, u, v int
  type Road struct {
	   u int
	   v int
  }
  var roads []Road
  fmt.Scanf("%v\n", &q)
  if q >= 1 && q <= 10 {
    fmt.Scanf("%v %v %v %v\n", &n, &m, &clib, &croad)
    if n >= 1 && n <= 10e+5 {
      if m >= 0 && m <= min(10e+5, n*(n-1)/2) {
        if croad >= 1 && clib >= 1 && croad <= 10e+5 && clib <= 10e+5 {
          roads = make([]Road, m)
          for i := range roads {
            fmt.Scanf("%v %v\n", &u, &v)
            if u >= 1 && u <= n && v >= 1 && v <= n {
              roads[i] = Road{u, v}
            }
          }
          fmt.Println(q, n, m, clib, croad)
          fmt.Println(roads)
        }
      }
    }
  }
}
func min(a, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}
