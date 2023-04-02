package errors

type InvalidField struct {
	Name        string
	Description string
}

type ValidationErr struct {
	message       string
	invalidFields []InvalidField
}
