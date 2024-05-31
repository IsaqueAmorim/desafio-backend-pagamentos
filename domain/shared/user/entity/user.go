package user

import "github.com/google/uuid"

type User struct {
	ID         string
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	Password   string
}

func NewUser(firstName string, middleName string, lastName string, email string, password string) *User {
	return &User{
		ID:         uuid.New().String(),
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Email:      email,
		Password:   password,
	}
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) SetFirstName(firstName string) {
	u.FirstName = firstName
}

func (u *User) GetMiddleName() string {
	return u.MiddleName
}

func (u *User) SetMiddleName(middleName string) {
	u.MiddleName = middleName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) SetLastName(lastName string) {
	u.LastName = lastName
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}
