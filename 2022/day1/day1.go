package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed calorie-inventory.txt
var calorieInventory string

func main() {
	lines := strings.Split(calorieInventory, "\n")

	var mostCaloriesSoFar float64
	var caloriesForCurrentElf float64
	for _, line := range lines {
		if line == "" {
			mostCaloriesSoFar = math.Max(caloriesForCurrentElf, mostCaloriesSoFar)
			caloriesForCurrentElf = 0
			continue
		}

		caloriesInLineItem, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		caloriesForCurrentElf += float64(caloriesInLineItem)

	}

	fmt.Println("The elf with the most calories has", mostCaloriesSoFar)
}
