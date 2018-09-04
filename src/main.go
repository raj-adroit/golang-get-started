package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, inactive},
	}

	grid := PowerGrid{300, plants}

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Choose an option: ")
	if option, err := requestOption(); err == nil {

		switch option {
		case "1":
			grid.generatePlantReport()
		case "2":
			grid.generateGridReport()
		}
	} else {
		fmt.Println(err.Error())
	}

}

func generatePlanCapacity(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d Capacity %0.f\n", idx, cap)
	}
}

func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantID := range activePlants {
		capacity += plantCapacities[plantID]
	}

}

func requestOption() (option string, err error) {
	fmt.Scanln(&option)
	if option != "1" && option != "2" {
		err = errors.New("Invalid option")
	}
	return
}

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "InActive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for idx, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type", p.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity", p.capacity)
		fmt.Printf("%-20s%s\n", "Status", p.status)
		fmt.Println("")
	}
}

func (pg *PowerGrid) generateGridReport() {
	capacity := 0.
	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}

	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity", capacity)
	fmt.Printf("%-20s%.0f\n", "Load", pg.load)
	fmt.Printf("%-20s%.2f%%\n", "Untilization", pg.load/capacity*100)
	fmt.Println("")
}
