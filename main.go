package main

import "fmt"

func main() {
	name := "RGR"
	completedLessons := 3
	totalLessons := 10
	dailyStudyHours := 1.5
	isConsistent := true

	remainingLessons := totalLessons - completedLessons
	progressCheck := completedLessons >= 5

	fmt.Println("Go Learning Score")
	fmt.Println("-----------------")
	fmt.Println("Name:", name)
	fmt.Println("Completed Lessons:", completedLessons)
	fmt.Println("Total Lessons:", totalLessons)
	fmt.Println("Remaining Lessons:", remainingLessons)
	fmt.Printf("Daily Study Hours: %.1f\n", dailyStudyHours)
	fmt.Println("Consistent:", isConsistent)
	fmt.Println("Progress Check:", progressCheck)
}
