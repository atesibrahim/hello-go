package defines

import "fmt"

func collections() {

	// defines array
	var arr [3]int
	arr[0] = 1
	arr[1] = 3
	arr[2] = 4
	fmt.Println(arr)
	fmt.Println(arr[1])

	// defines array second way
	arr2 := [3]int{3, 5, 7}
	fmt.Println(arr2)

	// slice - this is dynamic size
	slice := []int{3, 45, 7}
	fmt.Println(slice)
	slice = append(slice, 36)
	fmt.Println(slice)
	s1 := slice[1:]
	s2 := slice[:2]
	s3 := slice[1:2]
	fmt.Println(s1, s2, s3)

	// map
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 34
	delete(m, "foo")
	fmt.Println(m)

	// structs
	type user2 struct {
		ID        int
		FirstName string
		LastName  string
	}

	u := user2{ID: 1,
		FirstName: "Ibrahim",
		LastName:  "Ates",
	}
	fmt.Println(u)
}
