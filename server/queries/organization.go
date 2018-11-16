package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"

	"github.com/pwang347/cs304/server/common"
)

// CreateOrganization creates a new organization
func CreateOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result              sql.Result
		response            = SQLResponse{}
		tx                  *sql.Tx
		name                string
		contactEmailAddress string
		tmpRowsAffected     int64
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if contactEmailAddress, err = common.GetRequiredParam(params, "contactEmailAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO Organization (name,createdTimestamp,contactEmailAddress) VALUES(?,NOW(),?);",
		name, contactEmailAddress); err != nil {
		tx.Rollback()
		return
	}
	if result, err = tx.Exec("INSERT INTO UserOrganizationPairs (organizationName,userEmailAddress) VALUES(?,?);",
		name, contactEmailAddress); err != nil {
		return
	}
	if tmpRowsAffected, err = result.RowsAffected(); err != nil {
		return
	}
	response.AffectedRows += tmpRowsAffected
	if err = tx.Commit(); err != nil {
		return
	}
	if tmpRowsAffected, err = result.RowsAffected(); err != nil {
		return
	}
	response.AffectedRows += tmpRowsAffected
	data, err = json.Marshal(response)
	return
}

// DeleteOrganization deletes an organization
func DeleteOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result   sql.Result
		response = SQLResponse{}
		tx       *sql.Tx
		name     string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM Organization WHERE name=?;", name); err != nil {
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

// AddUserToOrganization adds a user to an existing organization
func AddUserToOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
		userEmailAddress string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if userEmailAddress, err = common.GetRequiredParam(params, "userEmailAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO UserOrganizationPairs (organizationName,userEmailAddress) VALUES(?,?);", organizationName, userEmailAddress); err != nil {
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

// RemoveUserFromOrganization removes a user from an existing organization
func RemoveUserFromOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
		userEmailAddress string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if userEmailAddress, err = common.GetRequiredParam(params, "userEmailAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM UserOrganizationPairs WHERE organizationName=? AND userEmailAddress=?;",
		organizationName, userEmailAddress); err != nil {
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

// QueryUserOrganizations queries all organizations a user belongs to
func QueryUserOrganizations(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		userEmailAddress string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if userEmailAddress, err = common.GetRequiredParam(params, "userEmailAddress"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM UserOrganizationPairs "+
		"WHERE userEmailAddress = ?;", userEmailAddress); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}
