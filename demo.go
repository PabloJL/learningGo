package main

import "fmt"

func main() {
	carsPerHour := (float64(1105)) * 100 / 100
	carsPerMinute := carsPerHour / 60
	carsCount:= 37
	var cost uint
	if carsCount < 10 {
		cost = uint(carsCount * 10000)
	}else if carsCount / 10 == 0 {
		cost = uint((carsCount / 10) * 95000)
		
	}else {
		group:= (carsCount / 10) * 95000
		singles:= (carsCount % 10) * 10000
		cost = uint(group + singles)
		
	}

	fmt.Println(carsPerHour, int(carsPerMinute), cost)
}

