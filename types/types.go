package types

import "time"

type UserStore interface {
	GetUsersByEmail(email string) (*User, error)
	GetUsersByID(id int) (*User, error)
	CreateUser(u *User) error
}

type mockUserStore struct {
}

func GetUsersByEmail(email string) (*User, error) {
	return nil, nil
}

func GetUsersByID(id int) (*User, error) {
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
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
