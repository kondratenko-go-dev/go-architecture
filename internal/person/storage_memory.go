package person

type MemoryPersonStorage struct {
	persons []*Person
	nextID  int64
}

func NewMemoryStorage() *MemoryPersonStorage {
	return &MemoryPersonStorage{
		persons: make([]*Person, 0),
		nextID:  1,
	}
}

func (s *MemoryPersonStorage) Save(person *Person) error {
	person.id = s.nextID
	s.nextID++
	s.persons = append(s.persons, person)
	return nil
}

func (s *MemoryPersonStorage) GetAll() ([]*Person, error) {
	return s.persons, nil
}

func (s *MemoryPersonStorage) FindByID(id int64) (*Person, error) {
	for _, p := range s.persons {
		if p.id == id {
			return p, nil
		}
	}
	return nil, ErrNotFound
}

func (s *MemoryPersonStorage) Delete(id int64) error {
	for i, p := range s.persons {
		if p.id == id {
			s.persons = append(s.persons[:i], s.persons[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (s *MemoryPersonStorage) FindByName(name string) ([]*Person, error) {
	result := make([]*Person, 0)
	for _, p := range s.persons {
		if p.name == name {
			result = append(result, p)
		}
	}
	return result, nil
}

func (s *MemoryPersonStorage) Update(person *Person) error {
	for i, p := range s.persons {
		if p.id == person.id {
			s.persons[i] = person
			return nil
		}
	}
	return ErrNotFound
}
