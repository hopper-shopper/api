package filter

type (
	IntOperator int

	IntFilter struct {
		Value    int
		Operator IntOperator
	}
)

const (
	OpEq IntOperator = iota
	OpLt
	OpLe
	OpGt
	OpGe
)

func NewIntFilter(value int, operator IntOperator) IntFilter {
	return IntFilter{
		Value:    value,
		Operator: operator,
	}
}

func (filter IntFilter) Match(value interface{}) bool {
	val, ok := value.(int)
	if !ok {
		return false
	}

	switch filter.Operator {
	case OpEq:
		return val == filter.Value
	case OpLt:
		return val < filter.Value
	case OpLe:
		return val <= filter.Value
	case OpGt:
		return val > filter.Value
	case OpGe:
		return val >= filter.Value
	default:
		return false
	}
}
