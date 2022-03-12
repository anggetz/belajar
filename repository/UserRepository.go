package repository

import (
	"belajar/golang/cores"
	"belajar/golang/internal/model"
	"fmt"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {

	return &UserRepository{}
}

//Get get data user
func (s *UserRepository) Get() ([]model.User, error) {
	db := cores.NewDatabaseMysql()
	err := db.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("could not get data from database, %s", err.Error())
	}

	defer db.Close()

	users := []model.User{}

	db.DB.Find(&users, "")

	return users, nil
}

func (s *UserRepository) Insert(user *model.User) error {

	db := cores.NewDatabaseMysql()
	err := db.ConnectDB()
	if err != nil {
		return fmt.Errorf("could not get data from database, %s", err.Error())
	}

	defer db.Close()

	tx := db.DB.Create(&user)
	if tx.Error != nil {
		return fmt.Errorf("could not create data to database, %s", tx.Error.Error())
	}

	return nil

}

func (s *UserRepository) Update() {

}

func (s *UserRepository) Delete() {

}

func (s *UserRepository) FindOne() {

}
