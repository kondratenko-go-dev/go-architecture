package main

type PersonService struct {
	storage PersonStorage
}

func NewPersonService(storage PersonStorage) *PersonService {
	return &PersonService{storage}
}

func (s *PersonService) CreatePerson(name string, age int, city, street string) (*Person, error) {
	person, err := NewPerson(name, age, city, street)
	if err != nil {
		return nil, err
	}
	if err := s.storage.Save(person); err != nil {
		return nil, err
	}
	return person, nil
}

func (s *PersonService) ListPeople() ([]*Person, error) {
	return s.storage.GetAll()
}

func PrintGreeting(greeter []Greeter) {
	for _, g := range greeter {
		g.Greet()
	}
}
