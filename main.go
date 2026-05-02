package main

import "fmt"

func main() {
	name := "RGR"
	completedLessons := 6
	totalLessons := 10
	dailyStudyHours := 1.5
	isConsistent := true
	hasBuiltMiniProject := false

	remainingLessons := totalLessons - completedLessons
	progressPercentage := float64(completedLessons) / float64(totalLessons) * 100

	fmt.Println("Go Learning Score")
	fmt.Println("-----------------")
	fmt.Println("Name:", name)
	fmt.Println("Completed Lessons:", completedLessons)
	fmt.Println("Total Lessons:", totalLessons)
	fmt.Println("Remaining Lessons:", remainingLessons)
	fmt.Printf("Progress: %.1f%%\n", progressPercentage)
	fmt.Printf("Daily Study Hours: %.1f\n", dailyStudyHours)
	fmt.Println("Built Mini Project:", hasBuiltMiniProject)
	fmt.Println("Consistent:", isConsistent)

	fmt.Println()
	fmt.Println("Progress Message")
	fmt.Println("----------------")

	if completedLessons >= 5 && isConsistent && hasBuiltMiniProject {
		fmt.Println("You are ready for the next milestone")
	} else if completedLessons >= 5 && isConsistent {
		fmt.Println("Good progress. Build a mini project before moving forward")
	} else if completedLessons >= 5 {
		fmt.Println("You completed enough lessons, but consistency is important")
	} else {
		fmt.Println("Complete more lessons first")
	}
}
