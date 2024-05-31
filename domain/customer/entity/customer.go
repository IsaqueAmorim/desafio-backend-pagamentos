package entity

import (
	CPF "back-end-challenge/domain/customer/value-object/cpf"
	user "back-end-challenge/domain/shared/user/entity"
)

type Customer struct {
	user.User
	CPF    CPF.CPF
	Amount float64
}

func NewCustomer(firstName, middleName, lastName, cpf, email, password string) *Customer {
	c, err := CPF.NewCPF(cpf)

	if err != nil {
		return nil
	}

	return &Customer{
		User:   *user.NewUser(firstName, middleName, lastName, email, password),
		CPF:    *c,
		Amount: 0.00,
	}
}
