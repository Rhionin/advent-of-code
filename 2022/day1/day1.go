package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed calorie-inventory.txt
var calorieInventory string

func getCalorieCountForTopNElves(count int) float64 {
	inventory := parseInventory()

	var totalCalories float64
	for i := 0; i < count; i++ {
		totalCalories += inventory[i]
	}

	return totalCalories
}

func parseInventory() []float64 {
	inventory := []float64{}

	var caloriesForCurrentElf float64
	lines := strings.Split(calorieInventory, "\n")
	for _, line := range lines {
		if line == "" {
			inventory = append(inventory, caloriesForCurrentElf)
			caloriesForCurrentElf = 0
			continue
		}

		caloriesInLineItem, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		caloriesForCurrentElf += float64(caloriesInLineItem)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(inventory)))
	return inventory
}
