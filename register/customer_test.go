package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerIsEmailedWhenRegistrationIsSuccessful(t *testing.T) {

	customerRegister := NewCustomerRegister(spyEmailSender)

	customer, _ := customerRegister.registerCustomer("Mark", "Bradley", "mark@testing.com")

	assert.Equal(t, "Mark", customer.FirstName)
	assert.Equal(t, "Bradley", customer.LastName)
	assert.Equal(t, "mark@testing.com", customer.Email)

}

func TestCustomerIsNotEmailsWhenEmailAddressIsInvalid(t *testing.T) {

	customerRegister := NewCustomerRegister(spyEmailSender)

	_, err := customerRegister.registerCustomer("Mark", "Bradley", "NOT AN EMAIL")

	assert.Error(t, err)
}
