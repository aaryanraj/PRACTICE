package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"a":"apple","b":"ball"}`
	jsonStr1 := `{"fruits":{
					"a":"Apple",
					"b":"Banana"
					},
				"color":{
					"r":"red",
					"b":"blue"
					}
				}`
	var y map[string]interface{}
	x := make(map[string]string)

	json.Unmarshal([]byte(jsonStr), &x)
	json.Unmarshal([]byte(jsonStr1), &y)
	fmt.Println(x)
	fmt.Println(y)
}
