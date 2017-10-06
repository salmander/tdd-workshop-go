package set

func NewSet() *Set {
	return &Set{
		isEmpty: true,
		size:    0,
		values:  make([]string, 10),
	}
}

func NewSetWithLimit(limit int) *Set {
	return &Set{
		isEmpty: true,
		size:    0,
		values:  make([]string, limit),
	}
}

type Set struct {
	isEmpty bool
	size    int
	values  []string
}

func (s *Set) IsEmpty() bool {
	return s.isEmpty
}

func (s *Set) Add(value string) {
	if len(s.values) == s.size {
		temp := make([]string, len(s.values)+1)
		copy(temp, s.values)
		s.values = temp
	}

	for _, v := range s.values {
		if v == value {
			return
		}
	}

	s.values[s.size] = value
	s.isEmpty = false
	s.size++
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) Contains(value string) bool {
	for i := 0; i < s.size; i++ {
		if s.values[i] == value {
			return true
		}
	}

	return false
}

func (s *Set) Remove(value string) {
	for i := 0; i < s.size; i++ {
		if s.values[i] == value {
			s.values[i] = s.values[s.size-1]
			s.values[s.size] = ""
			s.size--

			if s.size == 0 {
				s.isEmpty = true
			}

			return
		}
	}
}
