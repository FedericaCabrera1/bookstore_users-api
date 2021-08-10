package mysql_utils

import (
	"strings"

	"github.com/FedericaCabrera1/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

//every time we have an error from the database we will be returning this func
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError) //attempt to convert the err we recieve to a pointer to MySQLError
	if !ok {                              //if we have an error when converting, then we validate the error string, if it contains that text, we return a not found error
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		} //if we were not able to convert the error, we return an error
		return errors.NewInternalServerError(("error parsing database response"))
	}
	/*if we have reached this point our err is of MySQLError format,
	so we can switch to the .Number and take the decision based on the number, if the case is number 1062, then return that error,
	with any other case return that other error */
	switch sqlErr.Number {
	case 1062: //meaning duplicated key
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError(("error processing request"))

}
