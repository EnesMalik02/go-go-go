package user

// Service handles business logic for Users
type Service struct {
	repo *Repository
}

// NewService creates a new instance of Service
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUsers() []User {
	return s.repo.GetAll()
}

func (s *Service) CreateUser(user User) User {
	return s.repo.Create(user)
}

func (s *Service) GetUser(id int) (*User, error) {
	return s.repo.GetByID(id)
}
