package licence

import (
	"strings"
	"time"
)

type Person struct {
	name                  string
	dateOfBirth           time.Time
	randomNumberGenerator RandomNumberGenerator
}

type RandomNumberGenerator interface {
	// returns a string as number can be 0
	createRandomNumber(numberOfDigits int) string
}

func (p Person) getInitials() string {
	var initials string
	for _, name := range strings.Split(p.name, " ") {
		initials += string([]rune(name)[0])
	}

	return initials
}

func (p Person) getDob() string {
	return p.dateOfBirth.Format("20060102")
}

func (p *Person) LicenceNumber() string {
	initials := p.getInitials()
	dob := p.getDob()
	randomNumber := p.randomNumberGenerator.createRandomNumber(3)

	return initials + dob + randomNumber
}
