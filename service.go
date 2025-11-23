package main

type PersonService struct {
	storage PersonStorage
}

func NewPersonService(storage PersonStorage) *PersonService {
	return &PersonService{storage}
}

func (s *PersonService) CreatePerson(
	name string,
	age int,
	city, street string,
) (*Person, error) {

	person, err := NewPerson(name, age, city, street)
	if err != nil {
		return nil, err
	}
	if err := s.storage.Save(person); err != nil {
		return nil, err
	}
	return person, nil
}

func (s *PersonService) UpdatePerson(
	id int64,
	name string,
	age int,
	city, street string,
) (*Person, error) {

	existing, err := s.storage.FindByID(id)
	if err != nil {
		return nil, err
	}

	updated, err := NewPerson(name, age, city, street)
	if err != nil {
		return nil, err
	}

	updated.id = existing.id

	if err := s.storage.Update(updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *PersonService) ListPeople() ([]*Person, error) {
	return s.storage.GetAll()
}

func (s *PersonService) GetPerson(id int64) (*Person, error) {
	return s.storage.FindByID(id)
}

func (s *PersonService) DeletePerson(id int64) error {
	return s.storage.Delete(id)
}

func (s *PersonService) FindByName(name string) ([]*Person, error) {
	return s.storage.FindByName(name)
}

func PrintGreeting(greeter []Greeter) {
	for _, g := range greeter {
		g.Greet()
	}
}
