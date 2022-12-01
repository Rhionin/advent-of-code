package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed calorie-inventory.txt
var calorieInventory string

func getPart1Result() float64 {
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

	return mostCaloriesSoFar
}
