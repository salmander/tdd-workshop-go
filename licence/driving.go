package licence

import "time"

type Person struct {
	name        string
	dateOfBirth time.Time
}

type RandomNumberGenerator interface {
	// returns a string as number can be 0
	createRandomNumber(numberOfDigits int) string
}
