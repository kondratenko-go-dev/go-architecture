package main

type PersonStorage interface {
	Save(person *Person) error
	GetAll() ([]*Person, error)
	FindByID(id int64) (*Person, error)
	Delete(id int64) error
}

type Greeter interface {
	Greet()
}
