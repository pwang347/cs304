package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

// CreateOrganization creates a new organization
func CreateOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response            = SQLResponse{Rows: 0}
		tx                  *sql.Tx
		name                string
		createdTimestampStr string
		createdTimestamp    int
		contactEmailAddress string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if createdTimestampStr, err = common.GetRequiredParam(params, "createdTimestamp"); err != nil {
		return
	}
	if createdTimestamp, err = strconv.Atoi(createdTimestampStr); err != nil {
		return
	}
	if contactEmailAddress, err = common.GetRequiredParam(params, "contactEmailAddress"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO Organization (name,createdTimestamp,contactEmailAddress) VALUES(?,?,?);",
		name, createdTimestamp, contactEmailAddress); err != nil {
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

// DeleteOrganization deletes an organization
func DeleteOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response = SQLResponse{Rows: 0}
		tx       *sql.Tx
		name     string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if _, err = tx.Exec("DELETE FROM Organization WHERE name=?;", name); err != nil {
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

// AddUserToOrganization adds a user to an existing organization
func AddUserToOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{Rows: 0}
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
	if _, err = tx.Exec("INSERT INTO UserOrganizationPairs (organizationName,userEmailAddress) VALUES(?,?);", organizationName, userEmailAddress); err != nil {
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

// RemoveUserFromOrganization removes a user from an existing organization
func RemoveUserFromOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{Rows: 0}
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
	if _, err = tx.Exec("DELETE FROM UserOrganizationPairs WHERE organizationName=? AND userEmailAddress=?;",
		organizationName, userEmailAddress); err != nil {
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
