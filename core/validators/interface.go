package validators

type IValidator interface {
	Validate(expected interface{}, actual interface{}) error
}
