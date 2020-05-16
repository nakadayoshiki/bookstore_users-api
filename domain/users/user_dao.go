package users

import (
	"fmt"

	"github.com/nakadayoshiki/bookstore_users-api/utils/date"
	"github.com/nakadayoshiki/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

//Get user
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

// Save user
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}

	user.DateCreated = date.GetNowString()

	usersDB[user.ID] = user

	return nil
}
