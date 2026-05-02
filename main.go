package main

import (
	"fmt"
)

func main() {
	name := "RGR"
	stack := "MERN"
	language := "Go"
	experience := 2
	work_hours := 12.5
	consistant := true
	isLearning := true

	fmt.Println("Learning Profile")
	fmt.Println("----------------")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Current Stack: %s\n", stack)
	fmt.Printf("Experience: %d years\n", experience)
	fmt.Printf("Target Language: %s\n", language)
	fmt.Printf("Daily Study Hours: %f\n", work_hours)
	fmt.Printf("Consistant: %t\n", consistant)

	fmt.Println("Types")
	fmt.Println("----------------")
	fmt.Printf("name : %T \n", name)
	fmt.Printf("current stack: %T \n", stack)
	fmt.Printf("yearsOfExperience : %T \n", experience)
	fmt.Printf("targetLanguage : %T \n", language)
	fmt.Printf("dailyStudyHours : %T \n", work_hours)
	fmt.Printf("isConsistent : %T \n", isLearning)

}
