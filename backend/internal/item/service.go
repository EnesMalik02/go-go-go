package item

// Service handles business logic for Items
type Service struct {
	repo *Repository
}

// NewService creates a new instance of Service
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetItems(skip, limit int) []Item {
	return s.repo.GetAll(skip, limit)
}

func (s *Service) GetItem(id int) (*Item, error) {
	return s.repo.GetByID(id)
}

func (s *Service) CreateItem(item Item) Item {
	return s.repo.Create(item)
}

func (s *Service) UpdateItem(id int, item Item) (*Item, error) {
	return s.repo.Update(id, item)
}

func (s *Service) DeleteItem(id int) error {
	return s.repo.Delete(id)
}
