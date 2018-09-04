package main

import "fmt"

func main() {
	plantCapacities := []float64{30, 40, 40, 60, 70, 80}
	activePlants := []int{0, 1}
	gridLoad := 75.

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Choose an option: ")

	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		generatePlanCapacity(plantCapacities...)
	case "2":
		generatePowerGridReport(activePlants, plantCapacities, gridLoad)
	default:
		println("No such option!")
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
	fmt.Printf("%-20s%.0f\n", "Capacity", capacity)
	fmt.Printf("%-20s%.0f\n", "Grid Load", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization", gridLoad/capacity*100)
}
