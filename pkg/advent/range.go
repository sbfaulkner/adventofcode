package advent

// Range defines a min and max value
type Range struct{ Min, Max int }

// Includes returns a boolean indicating if the specified value is in the range
func (r Range) Includes(value int) bool {
	return r.Min <= value && value <= r.Max
}
