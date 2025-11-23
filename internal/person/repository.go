package person

type Storage interface {
	Save(person *Person) error
	GetAll() ([]*Person, error)
	FindByID(id int64) (*Person, error)
	Delete(id int64) error
	FindByName(name string) ([]*Person, error)
	Update(person *Person) error
}
