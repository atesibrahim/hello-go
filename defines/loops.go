package defines

func loops() {
	for i := 0; i < 5; i++ {
		println(i)
		if i == 3 {
			break
		}
	}

	// infinite loop
	var i int
	for {
		if i == 5 {
			break
		}
	}

	// loop over slice
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	// loop over map
	wellKnowPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnowPorts {
		println(k, v) // prints http: 80, https: 443
	}

	for k := range wellKnowPorts {
		println(k) // prints only keys as http, https
	}

	for _, v := range wellKnowPorts {
		println(v) // prints only values as 80, 443
	}

}
