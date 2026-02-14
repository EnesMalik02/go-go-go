package user

import "errors"

// Repository handles data access for Users
type Repository struct {
	users []User
}

// NewRepository creates a new instance of Repository
func NewRepository() *Repository {
	return &Repository{
		users: []User{
			{ID: 1, Username: "jdoe", Email: "john@example.com"},
			{ID: 2, Username: "asmith", Email: "alice@example.com"},
		},
	}
}

func (r *Repository) GetAll() []User {
	return r.users
}

func (r *Repository) Create(user User) User {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return user
}

func (r *Repository) GetByID(id int) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
