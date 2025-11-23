package main

import (
	"fmt"

	"github.com/kondratenko-go-dev/go-architecture/internal/person"
)

func main() {

	storage := person.NewMemoryStorage()
	service := person.NewService(storage)

	user, err := readLine("How many users need to be added: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	convUser, err := strConv(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := 0; i < convUser; i++ {
		name, err := readLine("What is your name? ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		age, err := readLine("What is your age? ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		convAge, err := strConv(age)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		city, err := readLine("What is your city? ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		street, err := readLine("What is your street? ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		_, err = service.CreatePerson(name, convAge, city, street)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	people, err := service.ListPeople()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	greeters := make([]Greeter, 0, len(people))

	for _, p := range people {
		greeters = append(greeters, p)
	}

	PrintGreeting(greeters)
}
