package response

import (
	"errors"
	"net/http"
)

// error general
var (
	ErrNotFound       = errors.New("error not found")
	ErrUnAuthorized   = errors.New("unauthorized")
	ErrForbiddenAcces = errors.New("forbidden access")
)

var (
	// auth
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum 4 character")
	ErrAuthIsNoExists        = errors.New("auth is not exists")
	ErrUsernameAlreadyUsed   = errors.New("username already used")
	ErrPasswordNotMatch      = errors.New("password not match")
	ErrUserIdInvalid         = errors.New("user invalid")
	ErrCodeInvalid           = errors.New("code unique is invalid")
	ErrUserNotFound          = errors.New("user not found")
	ErrTotalVoteInvalid      = errors.New("total vote exceed the limit")
	ErrFormatCSVInvalid      = errors.New("format csv is invalid")

	// products
	ErrProductRequired = errors.New("product is required")
	ErrProductInvalid  = errors.New("product must have minimun 4 character")
	ErrStockInvalid    = errors.New("stock must be greater than 0")
	ErrPriceInvalid    = errors.New("price must be greater than 0")

	// transactions
	ErrAmountInvalid          = errors.New("amount invalid")
	ErrAmountGreaterThanStock = errors.New("amount greater than stock")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral         = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest      = NewError("bad request", "40000", http.StatusBadRequest)
	ErrorNotFound        = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
	ErrorUnAuthorized    = NewError(ErrUnAuthorized.Error(), "40102", http.StatusUnauthorized)
	ErrorForbiddenAccess = NewError(ErrForbiddenAcces.Error(), "40301", http.StatusForbidden)
)

var (
	// Error bad request
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorProductRequired       = NewError(ErrProductRequired.Error(), "40005", http.StatusBadRequest)
	ErrorProductInvalid        = NewError(ErrProductInvalid.Error(), "40006", http.StatusBadRequest)
	ErrorStockInvalid          = NewError(ErrStockInvalid.Error(), "40007", http.StatusBadRequest)
	ErrorPriceInvalid          = NewError(ErrPriceInvalid.Error(), "40008", http.StatusBadRequest)
	ErrorAmountInvalid         = NewError(ErrAmountInvalid.Error(), "40009", http.StatusBadRequest)
	ErrorCodeInvalid           = NewError(ErrCodeInvalid.Error(), "40010", http.StatusBadRequest)
	ErrorTotalVoteInvalid      = NewError(ErrTotalVoteInvalid.Error(), "40011", http.StatusBadRequest)
	ErrorFormatCSVInvalid      = NewError(ErrFormatCSVInvalid.Error(), "40012", http.StatusBadRequest)

	ErrorAuthIsNoExists   = NewError(ErrAuthIsNoExists.Error(), "40401", http.StatusNotFound)
	ErrorEmailAlreadyUsed = NewError(ErrUsernameAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():              ErrorNotFound,
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrProductInvalid.Error():        ErrorProductInvalid,
		ErrStockInvalid.Error():          ErrorStockInvalid,
		ErrPriceInvalid.Error():          ErrorPriceInvalid,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrAuthIsNoExists.Error():        ErrorAuthIsNoExists,
		ErrUsernameAlreadyUsed.Error():   ErrorEmailAlreadyUsed,
		ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,
		ErrUnAuthorized.Error():          ErrorUnAuthorized,
		ErrForbiddenAcces.Error():        ErrorForbiddenAccess,
		ErrCodeInvalid.Error():           ErrorCodeInvalid,
		ErrTotalVoteInvalid.Error():      ErrorTotalVoteInvalid,
		ErrFormatCSVInvalid.Error():      ErrorFormatCSVInvalid,
	}
)
