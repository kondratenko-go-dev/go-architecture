package person

import (
	"errors"
	"fmt"
)

type Address struct {
	City   string
	Street string
}

type Person struct {
	id      int64
	name    string
	age     int
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

func (p *Person) ID() int64        { return p.id }
func (p *Person) Name() string     { return p.name }
func (p *Person) Age() int         { return p.age }
func (p *Person) City() string     { return p.address.City }
func (p *Person) Street() string   { return p.address.Street }
func (p *Person) Address() Address { return p.address }

func (p *Person) Greet() {
	fmt.Printf("\n[%d] My name is %s, I am %d years old and I live in %s, %s.\n",
		p.id, p.name, p.age, p.address.City, p.address.Street,
	)
}
