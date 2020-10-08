package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "São Paulo", "SP", "Rio de Janeiro", "RJ":
		region, continent = "Sudeste", "America do Sul"
	case "Los Angeles", "LA", "S":
		region, continent = "California", "America do Norte"

	default:
		region, continent = "Unknown", "Unknown"
	}

	return region, continent
}

func main() {
	region, continent := location("São Paulo")
	fmt.Println(region, continent)
}
