package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dataMap := map[string]interface{}{
		"fruits": map[string]string{
			"a": "Apple",
			"b": "Banana",
		},
		"cars": map[string]string{
			"f": "Ferrari",
			"l": "Lamborgini",
		},
	}

	resp, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(resp))
	}
}
