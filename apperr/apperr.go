package apperr

import (
	"errors"
	"net/http"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (c *CustomError) Error() string {
	return c.Message
}

func (c *CustomError) ConvertToErrorResponse() ErrorResponse {
	return ErrorResponse{
		Message: c.Message,
	}
}

var (
	ErrDatabaseConnection = errors.New("cannot make connection pool")

	ErrScanUser          = NewCustomError(http.StatusInternalServerError, "failed to scan to user")
	ErrFindUsersQuery    = NewCustomError(http.StatusInternalServerError, "find user query error")
	ErrFindUserByIDQuery = NewCustomError(http.StatusInternalServerError, "find user by id query error")
	ErrFindUserByEmail   = NewCustomError(http.StatusInternalServerError, "find user by email query error")
	ErrNewUserQuery      = NewCustomError(http.StatusInternalServerError, "new user query error")
	ErrUserNotFound      = NewCustomError(http.StatusBadRequest, "user not found")
	ErrEmailALreadyUsed  = NewCustomError(http.StatusBadRequest, "email already used")
	ErrWrongCredentials  = NewCustomError(http.StatusUnauthorized, "wrong password or email")
	ErrUpdateUser        = NewCustomError(http.StatusInternalServerError, "user info is not updated")

	ErrNewVerifyRequest    = NewCustomError(http.StatusInternalServerError, "failed to create verification request")
	ErrDeleteVerifyRequest = NewCustomError(http.StatusInternalServerError, "query delete verification request")

	ErrEmailUnregistered = NewCustomError(http.StatusInternalServerError, "user email never registered")

	ErrRecordNotFound       = NewCustomError(http.StatusBadRequest, "record not found")
	ErrRecordDuplicateFound = NewCustomError(http.StatusBadRequest, "record duplicate found")
	ErrDatabaseQuery        = NewCustomError(http.StatusInternalServerError, "config query error")

	ErrGenerateHashPassword = NewCustomError(http.StatusInternalServerError, "couldn't generate hash password")
	ErrGenerateJWTToken     = NewCustomError(http.StatusInternalServerError, "can't generate jwt token")

	ErrTxCommit = NewCustomError(http.StatusInternalServerError, "commit transaction error")

	ErrInvalidBody         = NewCustomError(http.StatusBadRequest, "invalid body")
	ErrInvalidParameter    = NewCustomError(http.StatusBadRequest, "invalid parameter")
	ErrUnauthorized        = NewCustomError(http.StatusUnauthorized, "unauthorized")
	ErrBearerTokenNotExist = NewCustomError(http.StatusUnauthorized, "bearer token not exist")

	ErrNewResetToken         = NewCustomError(http.StatusInternalServerError, "failed to create reset token")
	ErrInvalidateResetToken  = NewCustomError(http.StatusInternalServerError, "failed to update token valid status")
	ErrFindTokenData         = NewCustomError(http.StatusInternalServerError, "error on find password reset token data")
	ErrTokenDataNotFound     = NewCustomError(http.StatusBadRequest, "token data not found")
	ErrTokenIsInvalid        = NewCustomError(http.StatusInternalServerError, "token is invalid")
	ErrTokenIsUsed           = NewCustomError(http.StatusInternalServerError, "token is already used")
	ErrResetPasswordNotFound = NewCustomError(http.StatusInternalServerError, "token for reset password not found")
	ErrInvalidSignature      = NewCustomError(http.StatusBadRequest, "invalid signature")

	ErrResetPassword = NewCustomError(http.StatusInternalServerError, "update password failed")
)
