package main

import "fmt"

func main() {

	// pointer
	a := 5
	b := &a
	*b = 10
	fmt.Println(a, *b)

	// Array
	// you must specify array size and type
	m_arr := [3]string{"Tim", "Tom", "Kim"}
	fmt.Println(m_arr)

	// Slice
	// Slice is Unlimited Array
	m_slice := []string{"Tim", "Tom", "Kim", "Jack"}
	m_slice = append(m_slice, "Park")
	fmt.Println(m_slice)

	// map
	// map[key]value
	m_map := map[string]string{"name": "Kim", "sex": "male", "city": "Seou"}
	for key, val := range m_map {
		fmt.Println(key, val)
	}

	// struct
}
