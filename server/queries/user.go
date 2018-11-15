package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

// CreateUser creates a new user
func CreateUser(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
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
	if _, err = tx.Exec("INSERT INTO User (emailAddress,firstName,lastName,passwordHash,isAdmin,twoFactorPhoneNumber) VALUES(?,?,?,?,?,?);",
		emailAddress, firstName, lastName, passwordHash, isAdmin, twoFactorPhoneNumber); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}

	response.AffectedRows = 1
	data, err = json.Marshal(response)
	return
}

// DeleteUser deletes an user
func DeleteUser(db *sql.DB, params url.Values) (data []byte, err error) {
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
	if _, err = tx.Exec("DELETE FROM User WHERE emailAddress=?;", emailAddress); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}

	response.AffectedRows = 1
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
