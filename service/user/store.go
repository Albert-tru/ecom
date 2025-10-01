package user

import (
	"database/sql"
	"fmt"

	"github.com/Albert-tru/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	//查询数据库,将查询的多行结果保存到row
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("未找到用户")
	}

	return u, nil
}

func (s *Store) CreateUser(user *types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.Firstname, user.Lastname, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	// 查询数据库
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("未找到用户")
	}

	return u, nil
}

func scanRowIntoUser(row *sql.Rows) (*types.User, error) {
	u := new(types.User)
	err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
