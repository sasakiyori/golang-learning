package main

func PointerFunc(x *int) {
	*x = 2
}

func A() {
	var a int = 1
	PointerFunc(&a)
}

func B() *int {
	var c int = 3
	return &c
}

func main() {
	// no memory escape
	A()

	// no memory escape
	var tmp int = 3
	PointerFunc(&tmp)

	// has memory escape
	PointerFunc(B())
}

// RUN `go build -gcflags=-m memory_escape.go` OR `go tool compile -m memory_escape.go`:
/*
# command-line-arguments
./memory_escape.go:3:6: can inline PointerFunc
./memory_escape.go:7:6: can inline A
./memory_escape.go:9:13: inlining call to PointerFunc
./memory_escape.go:12:6: can inline B
./memory_escape.go:17:6: can inline main
./memory_escape.go:19:3: inlining call to A
./memory_escape.go:19:3: inlining call to PointerFunc
./memory_escape.go:23:13: inlining call to PointerFunc
./memory_escape.go:26:15: inlining call to B
./memory_escape.go:26:13: inlining call to PointerFunc
./memory_escape.go:3:18: x does not escape
./memory_escape.go:13:6: moved to heap: c
*/
