package repositories

import (
	"postgresTakeWords/internal/models"

	"github.com/jmoiron/sqlx"
)

func InsertUser(db *sqlx.DB, user *models.User) error {
	err := db.QueryRow("insert into users (name, email, password) values ($1, $2, $3) returning id",
		user.Name, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByName(db *sqlx.DB, name string) (*models.User, error) {
	u := models.User{}
	err := db.Get(&u, "SELECT * FROM users WHERE name=$1", name)
	if err != nil {
		return &models.User{}, err
	}
	return &u, nil
}

func GetUserByID(db *sqlx.DB, userID int) (*models.User, error) {
	u := models.User{}
	err := db.Get(&u, "SELECT * FROM users WHERE id=$1", userID)
	if err != nil {
		return &models.User{}, err
	}

	return &u, nil
}
