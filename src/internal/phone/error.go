package phone


type ValidationError struct {
	Err string
}

func (v *ValidationError) Error() string {
	return v.Err
}

func NewValidationError(err string) error {
	return &ValidationError{Err: err}
}

var (
	ErrCouldNotFetchDataFromDatabase   = NewValidationError("could not fetch data from database")
)
