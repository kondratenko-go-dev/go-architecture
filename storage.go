package main

type MemoryPersonStorage struct {
	persons []*Person
	nextID  int64
}

func NewMemoryPersonStorage() *MemoryPersonStorage {
	return &MemoryPersonStorage{
		persons: make([]*Person, 0),
		nextID:  1,
	}
}

func (s *MemoryPersonStorage) GetAll() ([]*Person, error) {
	return s.persons, nil
}

func (s *MemoryPersonStorage) Save(person *Person) error {
	person.id = s.nextID
	s.nextID++
	s.persons = append(s.persons, person)
	return nil
}
