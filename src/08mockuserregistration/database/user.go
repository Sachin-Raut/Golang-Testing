package database

import (
	"database/sql"

	"github.com/Sachin-Raut/Golang-Testing/src/mockuserregistration/entity"
)

//UserRepository is
type UserRepository struct {
	db *sql.DB
}

//User is
type User interface {
	Add(user entity.User) error
}

//NewUserRepository is
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

//Add is
func (ur *UserRepository) Add(user entity.User) error {
	sql := "insert into user(name,email,description) values (?,?,?)"
	_, err := ur.db.Exec(sql, user.Name, user.Email, user.Description)
	if err != nil {
		return err
	}
	return nil
}
