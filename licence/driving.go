package licence

import (
	"strings"
	"time"
)

type Person struct {
	name        string
	dateOfBirth time.Time
}

type RandomNumberGenerator interface {
	// returns a string as number can be 0
	createRandomNumber(numberOfDigits int) string
}

type LicenseNumberGenerator struct {
	p            Person
	randomNumber RandomNumberGenerator
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

func (license LicenseNumberGenerator) GenerateLicenseNumber() string {
	return license.p.getInitials() + license.p.getDob() + license.randomNumber.createRandomNumber(3)
}
