package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"

	"github.com/pwang347/cs304/server/common"
)

// CreateAccessGroup creates a new access group
func CreateAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{Rows: 0}
		tx               *sql.Tx
		name             string
		organizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO AccessGroup (name,organizationName) VALUES(?,?);",
		name, organizationName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	if response.Rows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// DeleteAccessGroup deletes an organization
func DeleteAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{Rows: 0}
		tx               *sql.Tx
		name             string
		organizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM AccessGroup WHERE name=? AND organizationName=?;", name, organizationName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	if response.Rows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// AddUserToAccessGroup adds a user to an existing organization
func AddUserToAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                      sql.Result
		response                    = SQLResponse{Rows: 0}
		tx                          *sql.Tx
		accessGroupName             string
		accessGroupOrganizationName string
		userEmailAddress            string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if accessGroupName, err = common.GetRequiredParam(params, "accessGroupName"); err != nil {
		return
	}
	if accessGroupOrganizationName, err = common.GetRequiredParam(params, "accessGroupOrganizationName"); err != nil {
		return
	}
	if userEmailAddress, err = common.GetRequiredParam(params, "userEmailAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO UserAccessGroupPairs (accessGroupName,accessGroupOrganizationName,userEmailAddress) VALUES(?,?,?);",
		accessGroupName, accessGroupOrganizationName, userEmailAddress); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	if response.Rows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// RemoveUserFromAccessGroup removes a user from an existing organization
func RemoveUserFromAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                      sql.Result
		response                    = SQLResponse{Rows: 0}
		tx                          *sql.Tx
		accessGroupName             string
		accessGroupOrganizationName string
		userEmailAddress            string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if accessGroupName, err = common.GetRequiredParam(params, "accessGroupName"); err != nil {
		return
	}
	if accessGroupOrganizationName, err = common.GetRequiredParam(params, "accessGroupOrganizationName"); err != nil {
		return
	}
	if userEmailAddress, err = common.GetRequiredParam(params, "userEmailAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM UserAccessGroupPairs WHERE accessGroupName=? AND accessGroupOrganizationName=? AND userEmailAddress=?;",
		accessGroupName, accessGroupOrganizationName, userEmailAddress); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	if response.Rows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}
