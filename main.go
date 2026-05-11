package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	Salary    float64
	IsTopG    bool
	IsLoveGod bool
	Address   Address
}

type Address struct {
	District string
	City     string
}

func main() {

	persons := map[int]Person{
		1: {Name: "R G R", Age: 26, Salary: 1150, IsTopG: true, IsLoveGod: true, Address: Address{District: "Central", City: "Metropolis"}},
		2: {Name: "U J", Age: 25, Salary: 2300, IsTopG: false, IsLoveGod: false, Address: Address{District: "West", City: "Gotham"}},
	}

	topG := persons[1]
	topG.Salary = 1500
	persons[1] = topG

	// fmt.Println(persons[1])

	// john := Person{
	// 	IsTopG: true,
	// }

	trisky := Person{"R G R", 26, 1350, true, true, Address{"western", "piliyandala"}}

	doe := trisky
	doe.Name = "John Doe"

	promoteEmployee(&trisky)
	trisky.addAllowance(50)
	fmt.Println(trisky.City)

}

func printProfile(profile map[string]string) {
	for key, value := range profile {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func promoteEmployee(employee *Person) {
	employee.Salary += 100
}

func (employee *Person) addAllowance(amount int) {
	employee.Salary += float64(amount)
	fmt.Printf("Salary after allowance: %.2f\n", employee.Salary)
}
