package register

import (
	"errors"
	"regexp"
)

type Customer struct {
	FirstName string
	LastName  string
	Email     string
}

type EmailSender interface {
	SendEmail(toEmail, toName, subject, message string)
}

func NewCustomerRegister(emailSender EmailSender) CustomerRegister {
	return CustomerRegister{emailSender}
}

type CustomerRegister struct {
	emailSender EmailSender
}

func (cr CustomerRegister) registerCustomer(firstName, lastName, email string) (Customer, error) {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !Re.MatchString(email) {
		return Customer{}, errors.New("Invalid email")
	}

	customer := Customer{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	cr.emailSender.SendEmail(email, firstName, "Welcome", "You have registered")

	return customer, nil
}
