package apperror

import (
	"github.com/rajankumar549/glHackathon/src/interfaces/apperror"
)

type Error struct {
	code     string
	message  string
	httpCode int
}

func BadError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 400,
	}
}

func NotFoundError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 404,
	}
}

func InternalServerError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 500,
	}
}

func ForbiddenError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 403,
	}
}

func NotImplementedError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 501,
	}
}

func BadGatewayError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 502,
	}
}

func (err *Error) Error() string {
	return err.message
}

func (err *Error) Code() string {
	return err.code
}

func (err *Error) HttpCode() int {
	if err.httpCode < 200 {
		err.httpCode = 200
	}
	return err.httpCode
}

func (r repo) ErrorHandler(err error) (string, string, int) {
	appErr, ok := err.(*Error)

	if !ok {
		appErr = InternalServerError(err.Error(), "Something went wrong, please try after sometime")
	}

	return appErr.Code(), appErr.Error(), appErr.HttpCode()
}

type repo struct {
}

func Init() apperror.AppError {
	return repo{}
}
