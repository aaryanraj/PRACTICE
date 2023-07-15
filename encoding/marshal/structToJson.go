package main

import (
	"encoding/json"
	"fmt"
)

type CarType struct {
	Type       string `json:"type"`
	TrunkSpace bool   `json:"trunk-space"`
	Seats      int    `json:"seats"`
}

type Car struct {
	Name    string  `json:"name"`
	Brand   string  `json:"brand"`
	CarType CarType `json:"car-type"`
	Milage  int     `json:"milage"`
}

func main() {
	var c Car

	c = Car{
		Name:  "Desire",
		Brand: "Maruti Suzuki",
		CarType: CarType{
			Type:       "Sedan",
			TrunkSpace: true,
			Seats:      5,
		},
		Milage: 12,
	}

	resp, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(resp))
	}
}
