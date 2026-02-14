package message

type Repository struct {
	messages []Message
}

func NewRepository() *Repository {
	return &Repository{
		messages: []Message{
			{ID: 1, Content: "Hello World", Sender: "System"},
			{ID: 2, Content: "Welcome to FSD", Sender: "Admin"},
		},
	}
}

func (r *Repository) GetMessages() []Message {
	return r.messages
}

func (r *Repository) CreateMessage(msg Message) Message {
	msg.ID = len(r.messages) + 1
	r.messages = append(r.messages, msg)
	return msg
}

func (r *Repository) UpdateMessage(id int, content string) *Message {
	for i, msg := range r.messages {
		if msg.ID == id {
			r.messages[i].Content = content
			return &r.messages[i]
		}
	}
	return nil
}

func (r *Repository) DeleteMessage(id int) bool {
	for i, msg := range r.messages {
		if msg.ID == id {
			r.messages = append(r.messages[:i], r.messages[i+1:]...)
			return true
		}
	}
	return false
}
