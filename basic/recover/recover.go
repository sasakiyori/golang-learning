package main
import "fmt"

func PrintLine() {
	fmt.Printf("system is running normally\n")
}

func ProcessError() {
	if errmsg := recover(); errmsg != nil {
		fmt.Printf("get error message: %s\n", errmsg)
	}
}

// panic && recover, like try && catch in c++
func RecoverTest(index int) {
	defer ProcessError()

	a := [1]int{1}
	a[index] = 1
}

func PanicTest(index int) {
	b := [1]int{1}
	b[index] = 1
}

func main() {
	PrintLine()
	RecoverTest(2)	// system catches runtime error and then runs normally 
	PrintLine()
	PanicTest(2)	// system calls panic, exit
	PrintLine()
}