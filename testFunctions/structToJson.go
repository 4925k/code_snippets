package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	sunny := Weather{"Sunny", -10, map[string]bool{"BatBoat": true, "BatMobile": true, "BatCycle": true}}
	input2 := Condition{"Windy", map[string]float64{"Orbit 1": 14, "Orbit 2": 20, "Orbit 3": 12}}

	x, _ := json.Marshal(sunny)
	y, err := json.Marshal(input2.Weather)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sunny, string(x))
	fmt.Println(input2, string(y))
}

type city struct {
	Orbits   []Orbit
	Vehicles []BatVehicle
	Weathers []Weather
}

type Orbit struct {
	Name     string
	Distance float64 //miles
	Craters  float64
}

type BatVehicle struct {
	Name       string
	Speed      float64 //miles/hour
	CraterTime float64 //minutes
}

type Weather struct {
	Name             string
	CraterPercent    float64
	VehicleUsability map[string]bool
}

type Condition struct {
	Weather      string
	OrbitTraffic map[string]float64
}

type fastestRoute struct {
	path    string
	vehicle string
	time    float64
}
