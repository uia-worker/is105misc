package main
import "fmt"

func main() {
  
  setup_world()
  fmt.Prinln(database)
  display_world()

}

func setup_world() {
  // A procedure to initialize the world
  // empty the database
  []->database
  // add the initial facts to the database
  add([cities]) // a real "point" / location
  add([cost of library]) // a value
  add([cost of road])
  add(constraints) // every citizen should have access to a library
}

func display_world() {
  fmt.Println(cities, cost_library, cost_road)
  // 1 3 3 2 1
  // [{1 2} {3 1} {2 3}]
}
