package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEmailSender struct {
	mock.Mock
}

func (es MockEmailSender) SendEmail(toEmail, toName, subject, message string) {
	es.Called(toEmail, toName, subject, message)
}

func TestCustomerIsEmailedWhenRegistrationIsSuccessful(t *testing.T) {
	spyEmailSender := new(MockEmailSender)
	spyEmailSender.On("SendEmail", "mark@testing.com", "Mark", "Welcome", "You have registered")
	customerRegister := NewCustomerRegister(spyEmailSender)

	customer, err := customerRegister.registerCustomer("Mark", "Bradley", "mark@testing.com")

	assert.NoError(t, err)
	assert.Equal(t, "Mark", customer.FirstName)
	assert.Equal(t, "Bradley", customer.LastName)
	assert.Equal(t, "mark@testing.com", customer.Email)
	spyEmailSender.AssertExpectations(t)
}

func TestCustomerIsNotEmailsWhenEmailAddressIsInvalid(t *testing.T) {
	spyEmailSender := new(MockEmailSender)
	customerRegister := NewCustomerRegister(spyEmailSender)

	_, err := customerRegister.registerCustomer("Mark", "Bradley", "NOT AN EMAIL")

	assert.Error(t, err)
	spyEmailSender.AssertNotCalled(t, "SendEmail")
}
