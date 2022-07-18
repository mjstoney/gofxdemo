package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

func main() {
	jsonData := `[{"species": "duck", "description": "says QUACK"}, {"species": "eagle", "description": "predator"}]`
	var birds []Bird

	json.Unmarshal([]byte(jsonData), &birds)
	fmt.Printf("Birds : %v", birds)
}
