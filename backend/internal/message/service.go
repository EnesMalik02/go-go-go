package message

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetMessages() []Message {
	return s.repo.GetMessages()
}

func (s *Service) CreateMessage(msg Message) Message {
	return s.repo.CreateMessage(msg)
}

func (s *Service) UpdateMessage(id int, content string) *Message {
	return s.repo.UpdateMessage(id, content)
}

func (s *Service) DeleteMessage(id int) bool {
	return s.repo.DeleteMessage(id)
}
