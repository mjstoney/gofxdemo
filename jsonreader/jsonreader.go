package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TimeSeries struct {
	Times  []float32
	Values []float32
}
type DataPoint struct {
	time  float32
	value float32
}

func (dp *DataPoint) printDP() {
	fmt.Printf("Time: %f\tValue: %f\n", dp.time, dp.value)
}

func main() {
	//var jsondata string
	file, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	var rawTimeSeries *TimeSeries
	json.Unmarshal(file, &rawTimeSeries)

	//`[{"species": "duck", "description": "says QUACK"}, {"species": "eagle", "description": "predator"}]`

	//var data []DataPoint

	//json.Unmarshal([]byte(jsondata), &data)
	l := len(rawTimeSeries.Times)
	var data_arr []DataPoint
	for i := 0; i < l; i++ {
		data_arr = append(data_arr, (DataPoint{time: rawTimeSeries.Times[i], value: rawTimeSeries.Values[i]}))
	}

	for i := 0; i < len(data_arr); i++ {
		data_arr[i].printDP()
	}

	backtojson, err := json.Marshal(data_arr)
	fmt.Println(backtojson)
}
