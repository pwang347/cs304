package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/pwang347/cs304/server/common"
)

// CreateUser creates a new user
func CreateUser(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result               sql.Result
		response             = SQLResponse{}
		tx                   *sql.Tx
		emailAddress         string
		firstName            string
		lastName             string
		passwordHash         string
		isAdminStr           string
		isAdmin              bool
		twoFactorPhoneNumber string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if emailAddress, err = common.GetRequiredParam(params, "emailAddress"); err != nil {
		return
	}
	if firstName, err = common.GetRequiredParam(params, "firstName"); err != nil {
		return
	}
	if lastName, err = common.GetRequiredParam(params, "lastName"); err != nil {
		return
	}
	if passwordHash, err = common.GetRequiredParam(params, "passwordHash"); err != nil {
		return
	}
	if isAdminStr, err = common.GetRequiredParam(params, "isAdmin"); err != nil {
		return
	}
	if isAdmin, err = strconv.ParseBool(isAdminStr); err != nil {
		return
	}
	if twoFactorPhoneNumber, err = common.GetRequiredParam(params, "twoFactorPhoneNumber"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO User (emailAddress,firstName,lastName,passwordHash,isAdmin,twoFactorPhoneNumber) VALUES(?,?,?,?,?,?);",
		emailAddress, firstName, lastName, passwordHash, isAdmin, twoFactorPhoneNumber); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	if response.AffectedRows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// UpdateUser updates a user
func UpdateUser(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result               sql.Result
		response             = SQLResponse{}
		tx                   *sql.Tx
		emailAddress         string
		firstName            string
		lastName             string
		passwordHash         string
		isAdminStr           string
		isAdmin              bool
		twoFactorPhoneNumber string
		updateStatements		 []string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}

	// email address is the only truly required param
	if emailAddress, err = common.GetRequiredParam(params, "emailAddress"); err != nil {
		return
	}

	// At least one param must be specified for the update to be valid
	if firstName, err = common.GetRequiredParam(params, "firstName"); err != nil {
		append(updateStatements, "firstName = " + firstName)
	}
	if lastName, err = common.GetRequiredParam(params, "lastName"); err != nil {
		append(updateStatements, "lastName = " + lastName)
	}
	if passwordHash, err = common.GetRequiredParam(params, "passwordHash"); err != nil {
		append(updateStatements, "passwordHash = " + passwordHash)
	}
	if isAdminStr, err = common.GetRequiredParam(params, "isAdmin"); err != nil {
		if isAdmin, err = strconv.ParseBool(isAdminStr); err != nil {
			append(updateStatements, "isAdmin = " + isAdmin)
		}
		else {
			return
		}
	}
	if twoFactorPhoneNumber, err = common.GetRequiredParam(params, "twoFactorPhoneNumber"); err != nil {
		append(updateStatements, "isAdmin = " + isAdmin
	}

	if len(updateStatements) < 1 {
		return
	}

	if result, err = tx.Exec("UPDATE User SET " + strings.Join(updateStatements, ", ") + " WHERE emailAddress = ?;",
		emailAddress); err != nil {
		tx.Rollback()
	}
	if err = tx.Commit(); err != nil {
		return
	}
	if response.AffectedRows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// DeleteUser deletes an user
func DeleteUser(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result       sql.Result
		response     = SQLResponse{}
		tx           *sql.Tx
		emailAddress string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if emailAddress, err = common.GetRequiredParam(params, "emailAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM User WHERE emailAddress=?;", emailAddress); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	if response.AffectedRows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// SelectUser queries the user table
func SelectUser(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response     = SQLResponse{}
		tx           *sql.Tx
		emailAddress string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if emailAddress, err = common.GetRequiredParam(params, "emailAddress"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM User"+
		"WHERE emailAddress=?;", emailAddress); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// UserLogin queries the user table using a password
func UserLogin(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response     = SQLResponse{}
		tx           *sql.Tx
		emailAddress string
		passwordHash string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if emailAddress, err = common.GetRequiredParam(params, "emailAddress"); err != nil {
		return
	}
	if passwordHash, err = common.GetRequiredParam(params, "passwordHash"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM User "+
		"WHERE emailAddress=? AND passwordHash=?;", emailAddress, passwordHash); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}
