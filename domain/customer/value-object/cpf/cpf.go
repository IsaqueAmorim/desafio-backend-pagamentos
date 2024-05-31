package cpf

import (
	"errors"
	"strconv"
	"strings"
)

type CPF struct {
	Number string
}

func NewCPF(cpf string) (*CPF, error) {
	strings.ReplaceAll(cpf, ".", "")
	strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return nil, errors.New("um CPF deve conter 11 digitos")
	}

	return &CPF{
		Number: cpf,
	}, nil
}

func (c *CPF) GetNumber() string {
	return c.Number
}

func (c *CPF) GetCPF() string {
	return c.Number
}

func (c *CPF) GetFormatedCPF() string {
	return c.Number
}

func (c CPF) IsValid() bool {
	cpf := c.Number

	// Check if CPF is not made up of the same character repeated 11 times
	if cpf == strings.Repeat(string(cpf[0]), 11) {
		return false
	}

	// Calculate the first verification digit
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}
	result := sum % 11
	firstVerDigit := 0
	if result > 1 {
		firstVerDigit = 11 - result
	}

	// Check if the first verification digit matches the 10th character in the CPF
	if firstVerDigit != int(cpf[9]-'0') {
		return false
	}

	// Calculate the second verification digit
	sum = 0
	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (11 - i)
	}
	result = sum % 11
	secondVerDigit := 0
	if result > 1 {
		secondVerDigit = 11 - result
	}

	// Check if the second verification digit matches the 11th character in the CPF
	if secondVerDigit != int(cpf[10]-'0') {
		return false
	}

	return true
}
