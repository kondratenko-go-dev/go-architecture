package main

import (
	"errors"
	"fmt"
)

type Address struct {
	City, Street string
}

type Person struct {
	id   int64
	name string
	age  int
	address Address
}

func NewPerson(name string, age int, city, street string) (*Person, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if age <= 0 {
		return nil, errors.New("age must be greater than zero")
	}

	if city == "" {
		return nil, errors.New("city cannot be empty")
	}

	if street == "" {
		return nil, errors.New("street cannot be empty")
	}

	return &Person{
		name: name,
		age:  age,
		address: Address{
			City:   city,
			Street: street,
		},
	}, nil
}

func (p *Person) Name() string { return p.name }

func (p *Person) SetName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	p.name = name
	return nil
}

func (p *Person) Age() int { return p.age }

func (p *Person) SetAge(age int) error {
	if age <= 0 {
		return errors.New("age is too low")
	}
	p.age = age
	return nil
}

func (p *Person) City() string { return p.addr.City }

func (p *Person) SetCity(city string) error {
	if city == "" {
		return errors.New("city is empty")
	}
	p.addr.City = city
	return nil
}

func (p *Person) Street() string { return p.addr.Street }

func (p *Person) SetStreet(street string) error {
	if street == "" {
		return errors.New("street is empty")
	}
	p.addr.Street = street
	return nil
}

func (p *Person) ID() int64 { return p.id }

func (p *Person) Greet() {
	fmt.Printf("\nMy name is %s, I am %d years old and I live in %s, %s.\n",
		p.name, p.age, p.addr.City, p.addr.Street,
	)
}
