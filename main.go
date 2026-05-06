package main

import "fmt"

func main() {

	fmt.Printf("The sum of 1, 2, 3, 4 and 5 is: %d\n", sum([]int{1, 3, 5, 8, 100}))

	fmt.Printf("subtraction of 5 and 3.41 is: %.2f \n", sub(5, 3.41))

	fmt.Println("Status of course completion: ", getStatus(3))

	remainingLessons, progressPercentage, progressMessage := calculateProgress(3, 5)
	fmt.Println("Remaining lessons:", remainingLessons)
	fmt.Println("Progress percentage:", progressPercentage)
	fmt.Println("Progress message:", progressMessage)
}

func sum(numbers []int) int {
	total := 0

	for _, value := range numbers {
		total = total + value
	}
	return total
}

func sub(num1 int, num2 float64) float64 {
	return float64(num1) - num2
}

func getStatus(completedLessons int) string {
	if completedLessons >= 5 {
		return "Congratulations! You have completed the course."
	}

	return "Keep going! You have more lessons to complete."
}

func calculateProgress(completedLessons int, totalLessons int) (int, int, string) {
	remainingLessons := totalLessons - completedLessons
	progressPercentage := completedLessons * 100 / totalLessons
	return remainingLessons, progressPercentage, fmt.Sprintf("You have completed %d%% of the course.", progressPercentage)
}
