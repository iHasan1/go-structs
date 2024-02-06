package main

/*
structs are similar to classes in other languages.
They allow to group together related data and functions that can be passed as a single unit.
*/

import (
	"fmt"	
	"time"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

// func (u user) outputUserDetails() ----> (u user) is the receiver
func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

// if a pointer is not used as argument to the mutation method, then a copy is made and that is mutated instead of the original, so it is necessary to use a pointer to modify the original.
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Convention: a function used to create the struct object.
func newUser(userFirstName string, userLastName string, userBirthdate string) *user {
	return &user{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user
	
	// struct literal notation
	appUser = newUser(userFirstName, userLastName, userBirthdate)

	// outputUserDetails(&appUser)
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}