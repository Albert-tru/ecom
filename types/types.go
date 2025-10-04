package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(u *User) error
}

type mockUserStore struct {
}

func GetUserByEmail(email string) (*User, error) {
	return nil, nil
}

func GetUserByID(id int) (*User, error) {
	return nil, nil
}

func CreateUser(u *User) error {
	return nil
}

type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type LoginrUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ProductStore interface {
	GetProducts() ([]Product, error)
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageUrl"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
}
