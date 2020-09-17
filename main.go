package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// parsing a simple json object
	sampleJson := `{"FirstName": "Sai", "LastName": "Bagam"}`

	test := make(map[string]interface{})

	json.Unmarshal([]byte(sampleJson), &test)

	fmt.Println("first element in the array is ", test["FirstName"])

	// Parsing an array of json
	arrayJson := `[{"FirstName": "John", "LastName": "Doe"}, {"FirstName": "Foo", "LastName": "Bar"}]`

	var testMap []map[string]interface{}

	json.Unmarshal([]byte(arrayJson), &testMap)

	fmt.Println("resultant map is ", testMap)

	for _, v := range testMap {
		fmt.Println("FirstNames in a json array is ", v["FirstName"])

	}

	// Object inside an object
	sample := `{
				"FirstName": "Sai", 
				"LastName": "Bagam", 
				"Address": {
						"Street":"Phoenix"
				}
				}`

	dummyMap := make(map[string]interface{})

	json.Unmarshal([]byte(sample), &dummyMap)

	add := dummyMap["Address"].(map[string]interface{})

	fmt.Println("address is ", add["Street"])

}
