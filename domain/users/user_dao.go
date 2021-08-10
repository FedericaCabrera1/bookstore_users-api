package users

import (
	"fmt"

	"github.com/FedericaCabrera1/bookstore_users-api/datasources/mysql/users_db"
	"github.com/FedericaCabrera1/bookstore_users-api/utils/errors"
	"github.com/FedericaCabrera1/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);" //query takes 4 parameters
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr { //get a user from the database based on the user id
	//work with a pointer in order to make sure we are modifing the actual element and not the copy
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError((err.Error()))

	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil { /*as we are selecting all 5 columns (id, first name, etc) => we need to pass that in the Scan(),
		the parameters of the Scan() are the destinations, so we tell to take the id from the data base and populate the user.Id, etc t*/
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr { //as we are in the domain package, the user knows how to save itself, so we create a method, not a func
	/*way of creating a new record on the database:
	we are creating a new row by creating the statement, executing it with the parameters, taking the LastInsertId and updating the user */
	stmt, err := users_db.Client.Prepare(queryInsertUser) //we first prepare the statement, so we have the chance to check if its valid
	if err != nil {
		return errors.NewInternalServerError((err.Error()))
	}
	//if we reach this point, our query is valid and we can build it
	defer stmt.Close() //when creating a statement, its essential to close it
	//everything in the defer statement is going to be executed just before the function return
	/*result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated) we execute the query with all of the parameters that the query takes
	the above line is equivalent to the code from line 37 to 42, but is convenient to first prepare the statement and validate it than do it all together like in the line above*/

	//point where we are saving the user, so date created is now

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password) // we have our statement and we execute it to the database
	if saveErr != nil {                                                                                                         //if we have error it means that the email we are trying to save is already saved on the data base
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, error := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id) //we just care about the error
	//attempt to execute that statement based on first name, last, etc, where id is that
	if error != nil {
		return mysql_utils.ParseError(error)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError((err.Error()))
	}
	defer stmt.Close()

	if _, error := stmt.Exec(user.Id); error != nil {
		return mysql_utils.ParseError(error)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError((err.Error()))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status) //execute the query against the database
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close() //we put the defer when we know we have a valid result
	results := make([]User, 0)
	for rows.Next() { //(Next is a boolean that informs us there´s a next row)
		//iterate on each row and use them to populate the var user, if there´s no error, append it to the results
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil { //we pass a pointer in order to fill the original user not a copy
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 { //if the slice is empty return a not found error
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
