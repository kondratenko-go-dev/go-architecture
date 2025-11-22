package main

type PersonStorage interface {
	Save(person *Person) error
	GetAll() ([]*Person, error)
}

type Greeter interface {
	Greet()
}
