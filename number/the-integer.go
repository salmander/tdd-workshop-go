package number

import "errors"

type TheInteger struct {
	number int
}

func (i TheInteger) GetNative() int {
	return i.number
}

func (i TheInteger) IsEqualTo(number int) bool {
	return i.number == number
}

func (i TheInteger) IsGreaterThan(number int) bool {
	return i.number > number
}

func (i TheInteger) IsLessThan(number int) bool {
	return i.number > number
}

func (i TheInteger) Add(i2 TheInteger) TheInteger {
	return TheInteger{
		i.number + i2.number,
	}
}

func (i TheInteger) Subtract(i2 TheInteger) TheInteger {
	return TheInteger{
		i.number - i2.number,
	}
}

func (i TheInteger) Multiply(i2 TheInteger) TheInteger {
	return TheInteger{
		i.number * i2.number,
	}
}

func (i TheInteger) DivideBy(i2 TheInteger) (TheInteger, error) {
	if i2.number == 0 {
		return TheInteger{}, errors.New("division by zero")
	}
	answer := TheInteger{
		i.number / i2.number,
	}
	return answer, nil
}
