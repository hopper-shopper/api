package filter

type (
	BoolFilter struct {
		Value bool
	}
)

func NewBoolFilter(value bool) BoolFilter {
	return BoolFilter{
		Value: value,
	}
}

func (filter BoolFilter) Match(value interface{}) bool {
	val, ok := value.(bool)
	if !ok {
		return false
	}

	return val == filter.Value
}
