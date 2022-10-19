package helper

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	"net/http"
)

func GetErrorStatusCode(err error) int {
	if err.Error() == ErrorEmailAlreadyExists.Error() {
		return http.StatusConflict
	}

	if err.Error() == ErrorUsernameAlreadyExists.Error() {
		return http.StatusConflict
	}

	if err.Error() == ErrorInvalidLogin.Error() {
		return http.StatusBadRequest
	}

	if isPgErrorUniqueViolation(err) {
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}

func isPgErrorUniqueViolation(err error) bool {
	pgError, _ := err.(*pgconn.PgError)
	if pgError != nil {
		return pgError.Code == pgerrcode.UniqueViolation
	}
	return false
}
