package user

import (
	"errors"
	"strings"

	"github.com/Sachin-Raut/Golang-Testing/src/mockuserregistration/database"
	"github.com/Sachin-Raut/Golang-Testing/src/mockuserregistration/entity"
)

//Service is
type Service struct {
	userRepository     database.User
	badWordsRepository database.BadWords
}

//NewUserService is
func NewUserService(userRepository database.User, badWordsRepository database.BadWords) *Service {
	return &Service{
		userRepository:     userRepository,
		badWordsRepository: badWordsRepository,
	}
}

//Register is
func (c *Service) Register(user entity.User) error {
	badWords, err := c.badWordsRepository.FindAll()
	if err != nil {
		return err
	}
	for _, badWord := range badWords {
		if strings.Contains(user.Description, badWord) {
			return errors.New("Bad word found")
		}
	}
	if err = c.userRepository.Add(user); err != nil {
		return err
	}
	return nil
}
