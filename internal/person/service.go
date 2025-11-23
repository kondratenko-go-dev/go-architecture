package person

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage}
}

func (s *Service) CreatePerson(
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

func (s *Service) ListPeople() ([]*Person, error) {
	return s.storage.GetAll()
}

func (s *Service) GetPerson(id int64) (*Person, error) {
	return s.storage.FindByID(id)
}

func (s *Service) DeletePerson(id int64) error {
	return s.storage.Delete(id)
}

func (s *Service) FindByName(name string) ([]*Person, error) {
	return s.storage.FindByName(name)
}

func (s *Service) UpdatePerson(
	id int64,
	name string,
	age int,
	city, street string,
) (*Person, error) {
	_, err := s.storage.FindByID(id)
	if err != nil {
		return nil, err
	}

	updated, err := NewPerson(name, age, city, street)
	if err != nil {
		return nil, err
	}

	updated.id = id

	if err := s.storage.Update(updated); err != nil {
		return nil, err
	}

	return updated, nil
}
