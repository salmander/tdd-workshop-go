package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	empty := NewSet()

	assert.True(t, empty.IsEmpty(), "Set was expected to be empty")

	one := NewSet()
	one.Add("one")

	assert.False(t, one.IsEmpty(), "Set was expected to have values.")
}

func TestCount(t *testing.T) {
	empty := NewSet()

	one := NewSet()
	one.Add("one")

	many := NewSet()
	many.Add("one")
	many.Add("two")

	assert.Equal(t, 0, empty.Size())

	assert.Equal(t, 1, one.Size())

	assert.NotZero(t, many.Size())
	assert.NotEqual(t, 1, many.Size())
}

func TestContains(t *testing.T) {
	empty := NewSet()

	one := NewSet()
	one.Add("one")

	assert.False(t, empty.Contains("one"))

	assert.True(t, one.Contains("one"))
}

func TestDuplicateValues(t *testing.T) {
	one := NewSet()
	one.Add("one")
	one.Add("one")

	assert.Equal(t, 1, one.size)
}

func TestRemoveValues(t *testing.T) {
	one := NewSet()
	one.Add("one")
	one.Remove("one")

	assert.Equal(t, 0, one.size)
	assert.True(t, one.IsEmpty())
	assert.False(t, one.Contains("one"))

	cereal := NewSet()
	cereal.Add("snap")
	cereal.Add("crackle")
	cereal.Add("pop")

	cereal.Remove("crackle")

	assert.Equal(t, 2, cereal.size)
	assert.True(t, cereal.Contains("snap"))
	assert.True(t, cereal.Contains("pop"))
	assert.False(t, cereal.Contains("crackle"))
}

func TestGrow(t *testing.T) {
	growable := NewSetWithLimit(1)
	growable.Add("growable")
	growable.Add("two")

	assert.Equal(t, 2, growable.Size())
}
