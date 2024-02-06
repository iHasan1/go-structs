package user

import (
	"errors"
	"fmt"	
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

// func (u user) outputUserDetails() ----> (u user) is the receiver
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

// if a pointer is not used as argument to the mutation method, then a copy is made and that is mutated instead of the original, so it is necessary to use a pointer to modify the original.
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email string, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "admin",
			lastName: "admin",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}
// Convention: a function used to create the struct object.
func New(userFirstName string, userLastName string, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}