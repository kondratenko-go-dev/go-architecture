package main

type MemoryPersonStorage struct {
	persons []*Person
}

func NewMemoryPersonStorage() *MemoryPersonStorage {
	return &MemoryPersonStorage{
		persons: make([]*Person, 0),
	}
}

func (s *MemoryPersonStorage) GetAll() ([]*Person, error) {
	return s.persons, nil
}

func (s *MemoryPersonStorage) Save(person *Person) error {
	s.persons = append(s.persons, person)
	return nil
}
