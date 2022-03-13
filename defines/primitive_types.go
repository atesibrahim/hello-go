package defines

import "fmt"

// const expression define as global
const (
	firstConst = iota
	secondConst
)

func primitive_types() {
	var i int = 42
	fmt.Println(i)

	firstName := "Ibrahim"
	fmt.Println(firstName)

	// define pointer
	var lastName *string = new(string)
	*lastName = "Ates"
	fmt.Println(*lastName)

	secondName := "Halil"

	// assign secondName to ptr
	ptr := &secondName
	fmt.Println(ptr, *ptr)

	secondName = "NoOne"
	fmt.Println(ptr, *ptr)

	// declare const as local
	const pi = 3.1415
	fmt.Println(pi)

	fmt.Println(firstConst, secondConst)

}
