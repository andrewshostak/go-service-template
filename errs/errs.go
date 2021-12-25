package errs

type Error interface {
	error
	GetCause() ErrorType
	GetError() error
}

type ErrorType int

const (
	UserError   ErrorType = 0
	ServerError ErrorType = 1
)

func New(err error, cause ErrorType) Error {
	return &errs{
		error: err,
		cause: cause,
	}
}

type errs struct {
	error error
	cause ErrorType
}

func (e errs) Error() string {
	return e.error.Error()
}

func (e errs) GetCause() ErrorType {
	return e.cause
}

func (e errs) GetError() error {
	return e.error
}