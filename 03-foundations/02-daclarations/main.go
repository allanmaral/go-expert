package main

const a = "Hello, World!"

var (
	b bool    = true
	c int     = 10
	d string  = "John"
	e float64 = 1.2
)

func main() {
	// a = "b" // cannot assign to a (untyped string constant "Hello, World!")
	// b = true

	f := "X"
	// f := "test" // no new variables on left side of :=
	f = "test"

	println(a) // Hello, World!
	println(b) // true
	println(c) // 10
	println(d) // John
	println(e) // +1.200000e+000
	println(f) // test
}
