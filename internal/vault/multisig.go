package vault

// MultiSigRequirement defines the threshold of signatures needed for an action.
type MultiSigRequirement struct {
	Threshold int
	Owners    []string
}

// MultiSigTransaction tracks the collection of signatures for a specific action.
type MultiSigTransaction struct {
	ActionID   string
	Signatures map[string][]byte
}

// CheckThreshold verifies if the transaction has reached the required number of valid signatures.
func (r *MultiSigRequirement) CheckThreshold(tx *MultiSigTransaction) bool {
	validCount := 0
	for _, owner := range r.Owners {
		if _, exists := tx.Signatures[owner]; exists {
			validCount++
		}
	}
	return validCount >= r.Threshold
}
