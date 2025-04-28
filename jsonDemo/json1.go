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

func jsonMarshalDemo() {
	pln("jsonMarshalDemo...")
	getJson(true)
	getJson(false)
	getJson(1)
	getJson(1.1)
	getJson("foobar")
	getJson([]string{"foo", "bar", "baz"})
	getJson(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	getJson(employee{"emp1", 1, 25, "addr1", "engg"})
	getJson(person{"p1", 1011, 45, "addr2", "accountant"})
	pln("-----------------------------------------------")
}

func getData(buf []byte, data any) {
	err := json.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}
	pln("json buf", string(buf))
	pln("unmarshalled data: ", data)
}

func jsonUnmarshalGeneric() {
	pln("jsonUnmarshalGeneric...")
	buf := []byte(`{ "foo1": "bar1", "foo2": 222, "foo3": 1.01, "foo4": ["bar41", "bar42"], "foo5": {"foo51": 51, "foo52": "bar52"} }`)
	var data map[string]interface{}
	getData(buf, &data)

	pln("foo1:", data["foo1"].(string))
	pln("foo2:", data["foo2"].(float64))
	pln("foo3:", data["foo3"].(float64))
	foo4 := data["foo4"].([]interface{})
	pln("foo4:", foo4)
	pln("foo4[0]:", foo4[0].(string))
	pln("foo4[1]:", foo4[1].(string))
	foo5 := data["foo5"].(map[string]interface{})
	pln("foo5:", foo5)
	pln("foo5[\"foo51\"]:", foo5["foo51"].(float64))
	pln("foo5[\"foo52\"]:", foo5["foo52"].(string))
	pln("---------")

}

func jsonUnmarshalTyped() {
	pln("jsonUnmarshalTyped...")
	buf := []byte(`{"name":"p1","id":1011,"age":45,"address":"addr2", "xxx": "yyy"}`)
	p := person{}
	getData(buf, &p)
	pln("name:", p.Name)
	pln("id:", p.Id)
	pln("age:", p.Age)
	pln("address:", p.Addr)
	pln("job:", p.Job)
	pln("---------")
}

func jsonUnmarshalDemo() {
	pln("jsonUnmarshalDemo...")
	jsonUnmarshalGeneric()
	jsonUnmarshalTyped()
	pln("-----------------------------------------------")
}

func jsonDemo() {
	jsonMarshalDemo()
	jsonUnmarshalDemo()
}

func main() {
	jsonDemo()
}
