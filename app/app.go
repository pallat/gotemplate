package app

const (
	storeErrorStutas = 450
)

type status string

const (
	Success status = "success"
	Fail    status = "error"
)

type Response struct {
	Status  status
	Message string
	Data    any
}

type Error Response

func (err *Error) Error() string {
	return err.Message
}
