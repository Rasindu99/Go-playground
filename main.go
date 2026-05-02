package main

import "fmt"

func main() {

	topics := []string{"Variables", "Data Types", "If Else", "Loops", "Functions"}
	var topicCount int

	fmt.Println("Go Study Planner \n----------------")
	for index, topic := range topics {
		fmt.Printf("%d. %s\n", index+1, topic)
	}

	fmt.Println("\nGo Study Planner \n----------------")
	for _, topic := range topics {
		topicCount++
		fmt.Printf("Studing: %s\n", topic)
	}
	fmt.Printf("\nTotal Topics: %d\n", topicCount)
}
