package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItReturnsPrimitiveValue(t *testing.T) {
	one := TheInteger{
		1,
	}

	assert.Equal(t, 1, one.GetNative())

	two := TheInteger{
		2,
	}

	assert.Equal(t, 2, two.GetNative())
}

func TestItCanComparePrimitiveValues(t *testing.T) {
	one := TheInteger{
		1,
	}

	assert.Equal(t, true, one.IsEqualTo(1))
	assert.Equal(t, false, one.IsEqualTo(2))
}

func TestItCanCompareNumbersGreaterThanItSelf(t *testing.T) {
	twentyOne := TheInteger{
		21,
	}

	assert.Equal(t, true, twentyOne.IsGreaterThan(20))
	assert.Equal(t, true, twentyOne.IsGreaterThan(19))
	assert.Equal(t, true, twentyOne.IsGreaterThan(18))
	assert.Equal(t, false, twentyOne.IsGreaterThan(21))
	assert.Equal(t, false, twentyOne.IsGreaterThan(22))
}

func TestItCanCompareNumbersLessThanItSelf(t *testing.T) {
	twentyOne := TheInteger{
		21,
	}

	assert.Equal(t, true, twentyOne.IsLessThan(20))
	assert.Equal(t, true, twentyOne.IsLessThan(19))
	assert.Equal(t, true, twentyOne.IsLessThan(18))
	assert.Equal(t, false, twentyOne.IsLessThan(21))
	assert.Equal(t, false, twentyOne.IsLessThan(22))
}

func TestItCanAddAnotherIntegerAndReturnsNewInteger(t *testing.T) {
	one := TheInteger{
		1,
	}
	two := TheInteger{
		2,
	}

	actual := one.Add(two)

	assert.Equal(t, 1, one.GetNative())
	assert.Equal(t, 2, two.GetNative())
	assert.Equal(t, 3, actual.GetNative())
}

func TestItCanSubtractAnotherIntegerAndReturnsNewInteger(t *testing.T) {
	one := TheInteger{
		1,
	}
	two := TheInteger{
		2,
	}

	actual := two.Subtract(one)

	assert.Equal(t, 1, one.GetNative())
	assert.Equal(t, 2, two.GetNative())
	assert.Equal(t, 1, actual.GetNative())
}

func TestItCanMultiplyAnotherIntegerAndReturnsNewInteger(t *testing.T) {
	three := TheInteger{
		3,
	}
	two := TheInteger{
		2,
	}

	actual := two.Multiply(three)

	assert.Equal(t, 3, three.GetNative())
	assert.Equal(t, 2, two.GetNative())
	assert.Equal(t, 6, actual.GetNative())
}

func TestItCanDivideAnotherIntegerAndReturnsNewInteger(t *testing.T) {
	four := TheInteger{
		4,
	}
	two := TheInteger{
		2,
	}

	actual, err := four.DivideBy(two)
	assert.NoError(t, err)

	assert.Equal(t, 4, four.GetNative())
	assert.Equal(t, 2, two.GetNative())
	assert.Equal(t, 2, actual.GetNative())
}

func TestItReturnsAnErrorWhenDoingDivisionByZero(t *testing.T) {
	five := TheInteger{
		5,
	}

	zero := TheInteger{
		0,
	}

	actual, err := five.DivideBy(zero)

	assert.Error(t, err)
	assert.Equal(t, 0, actual.GetNative())
}

func TestItDoesRoundingWithRemainders(t *testing.T) {
	five := TheInteger{
		5,
	}

	two := TheInteger{
		2,
	}

	actual, _ := five.DivideBy(two)

	assert.Equal(t, 2, actual.GetNative())
}
