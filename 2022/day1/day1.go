package main

import (
	"fmt"

	_ "embed"
)

//go:embed calorie-inventory.txt
var calorieInventory string

func main() {
	fmt.Println("calorieInventory:", calorieInventory)
}
