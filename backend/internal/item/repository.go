package item

import "errors"

// Repository handles data access for Items
type Repository struct {
	items []Item
}

// NewRepository creates a new instance of Repository
func NewRepository() *Repository {
	return &Repository{
		items: []Item{
			{ID: 1, Name: "Laptop", Price: 1500.00, IsOffer: true},
			{ID: 2, Name: "Mouse", Price: 25.50, IsOffer: false},
		},
	}
}

func (r *Repository) GetAll(skip, limit int) []Item {
	start := skip
	end := skip + limit

	if start > len(r.items) {
		start = len(r.items)
	}
	if end > len(r.items) {
		end = len(r.items)
	}
	return r.items[start:end]
}

func (r *Repository) GetByID(id int) (*Item, error) {
	for _, item := range r.items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *Repository) Create(item Item) Item {
	item.ID = len(r.items) + 1
	r.items = append(r.items, item)
	return item
}

func (r *Repository) Update(id int, updatedItem Item) (*Item, error) {
	for i, item := range r.items {
		if item.ID == id {
			updatedItem.ID = id
			r.items[i] = updatedItem
			return &updatedItem, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *Repository) Delete(id int) error {
	for i, item := range r.items {
		if item.ID == id {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}
