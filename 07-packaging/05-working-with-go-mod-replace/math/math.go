package math

// Exported variables and methods need to start with a Capital letter
var ExportedVariable = "exported"

// Variables and methods that starts with lower case letters are not visible outside
// their modules
var notExportedVariable = "not exported variable"

// Naming an struct with lower case first letter,
// makes the struct constructor private
type math struct {
	a          int
	b          int
	SomePublic string
}

func New(a, b int) math {
	return math{a: a, b: b, SomePublic: "value"}
}

func (m math) Add() int {
	return m.a + m.b
}
