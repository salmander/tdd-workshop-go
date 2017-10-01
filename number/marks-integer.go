package number

type MarksInteger interface {
	GetNative() int

	IsEqualTo(integer MarksInteger) bool

	IsGreaterThan(integer MarksInteger) bool

	IsLessThan(integer MarksInteger) bool

	Add(integer MarksInteger) MarksInteger

	Subtract(integer MarksInteger) MarksInteger

	Multiply(integer MarksInteger) MarksInteger

	DivideBy(integer MarksInteger) MarksInteger
}
