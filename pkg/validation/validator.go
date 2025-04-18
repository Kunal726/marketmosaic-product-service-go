package validation

// Validator interface defines the contract for validatable entities
type Validator interface {
	Validate() error
}
