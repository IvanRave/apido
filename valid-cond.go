package apido

// ValidCond
type ValidCond struct {
	UnMatched map[string]string `json:"conditions,omitempty"`
}

// IsValidated Quick property to identify errors instead len(array)
func (validCond *ValidCond) IsValidated() bool {
	return len(validCond.UnMatched) == 0
}
