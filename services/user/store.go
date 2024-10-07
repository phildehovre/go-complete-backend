package user

import (
	"database/sql"
	"fmt"

	"github.com/phildehovre/go-complete-backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {

	fmt.Println("User submitted to GetUserByEmail: ", email)
	rows, err := s.db.Query(`SELECT * FROM users WHERE email = ?`, email)
	if err != nil {
		return nil, err
	}
	var user = new(types.User)
	for rows.Next() {
		userRow, err := scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
		user = userRow
		defer rows.Close()
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found %s", email)
	}
	return user, nil
}

func (s *Store) CreateUser(user *types.User) error {
	fmt.Println("User submitted to CreateUser: ", user)
	_, err := s.db.Exec(`INSERT INTO users (firstName, lastName, email, password) VALUES (?,?,?,?)`,
		user.Firstname, user.Lastname, user.Email, user.Password)
	return err
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	var user = new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
