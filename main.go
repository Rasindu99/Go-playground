package main

import "fmt"

func printSkills(skills []string) {
	for index, skill := range skills {
		fmt.Printf("%d. %s\n", index+1, skill)
	}
}

func main() {

	user := map[string]string{
		"name":     "R G R",
		"age":      "26",
		"salary":   "1250$",
		"attitude": "Never Ever Give Up",
	}

	user["hobby"] = "Gaming"
	user["salary"] = "1370$"
	delete(user, "hobby")

	fmt.Println(user)

	for key, value := range user {
		fmt.Printf("%s : %s\n", key, value)
	}

	value, exists := user["hobby"]

	if exists {
		fmt.Println("Hobby:", value)
	} else {
		fmt.Println("Hobby key does not exist in the user map.")
	}
}
