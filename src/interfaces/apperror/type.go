package apperror

type AppError interface {
	ErrorHandler(error) (string, string, int)
}
