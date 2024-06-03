package validations

type Validate[T any] interface {
	Validate(option T)
	CanValidate(option T) bool
}
