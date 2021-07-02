package main
import "fmt"

// test interface && struct nesting
type HighSchoolStudent struct {
	Student
	ranking int
}

type Student struct {
	id int32
	name string
}

type ITest interface {
	Id() int32
	Name() string
}

func (this *Student) Id() int32 {
	return this.id
}

func (this *Student) Name() string {
	return this.name
}

func main() {
	var test = new(HighSchoolStudent)
	test.id = 1
	test.name = "bob"
	test.ranking = 1
	fmt.Printf("id: %d, name: %s\n", test.Id(), test.Name())
}
