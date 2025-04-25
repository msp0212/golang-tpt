package main

import (
	"encoding/json"
	"fmt"
)

var pln = fmt.Println
var pf = fmt.Printf

/*
Support for json encoding is already built into golang.
A field MUST an exported field (start with Capital letter)
to be encoded into json
*/
type employee struct {
	Name       string
	Id         uint
	Age        uint
	Addr       string
	Department string
}

/*
Json tags can be added for the field names.
These tags names will appear as key names in the generated json
Make sure there is no space after "json:"
*/
type person struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
	Age  uint   `json:"age"`
	Addr string `json:"address"`
	Job  string `json:"job"`
}

func getJson[T any](data T) []byte {
	buf, err := json.Marshal(data)
	if err != nil {
		pln("error in json marshal")
		return nil
	}
	pf("json of %T : %s\n", data, buf)
	return buf
}

func jsonDemo() {
	getJson(true)
	getJson(false)
	getJson(1)
	getJson(1.1)
	getJson("foobar")
	getJson([]string{"foo", "bar", "baz"})
	getJson(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	getJson(employee{"emp1", 1, 25, "addr1", "engg"})
	getJson(person{"p1", 1011, 45, "addr2", "accountant"})
}

func main() {
	jsonDemo()
}
