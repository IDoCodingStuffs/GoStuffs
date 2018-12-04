package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func location(city string) (region, continent string) {
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "CA", "North America"
	case "New York", "NY", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return
}

func main() {
	fmt.Println(add(4, 5))
	region, continent := location("LA")
	fmt.Printf("LA is in %s, %s", region, continent)
}
