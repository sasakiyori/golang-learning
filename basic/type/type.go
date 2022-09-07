package main
import "fmt"

// test multiple type-defining
type(
	name string
	id   int
)


func main() {
	var myname name = "test"
	var myid id = 1
	fmt.Printf("name: %s, id: %d\n", myname, myid)
}
