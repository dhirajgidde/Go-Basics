package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	jsonString := `{"name":"battery sensor","capacity":40,"time":"2019-01-2T19:07:28Z"}`
	fmt.Println(jsonString)

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("%+v\n", reading)
	}

}
