package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating an empty map
	var m1 map[string]int
	m1 = make(map[string]int)

	//set operation on key
	m1["key1"] = 1
	m1["key2"] = 2
	fmt.Println("m1 = ", m1)

	//get operation on key
	fmt.Println("key1 = ", m1["key1"], "key2 = ", m1["key2"])

	//get operation on a non existent key returns zero value for that type
	val, ifExists := m1["key3"]
	fmt.Println("get operation on non existent key, key3 = ", val, " ,ifExists = ", ifExists)

	//delete operation on key
	delete(m1, "key1")
	//test for existence of key without retrieveing the value by using an underscore _
	_, ifExists = m1["key1"]
	fmt.Println("After delete operation, key1 exists =  ", ifExists)

	//clear operation on map
	clear(m1)
	fmt.Println("m1 after clear operation= ", m1)

	m2 := make(map[string]int)
	fmt.Println("m2 = ", m2)
	//equal operation on map
	fmt.Println("m1 == m2 : ", maps.Equal(m1, m2))
}
