package handlerError

type NotfoundError struct {
	Message string
}
type InternalServerError struct {
	Message string
}
type UnauthorizedError struct {
	Message string
}
type BadRequestError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}
func (e *UnauthorizedError) Error() string {
	return e.Message
}
func (e *NotfoundError) Error() string {
	return e.Message
}
func (e *BadRequestError) Error() string {
	return e.Message
}
