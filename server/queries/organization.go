package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/cs304/server/common"
)

// CreateUser creates a new user
func CreateUser(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response             = SQLResponse{Rows: 0}
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

	response.Rows = 1
	data, err = json.Marshal(response)
	return
}
