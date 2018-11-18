package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/pwang347/cs304/server/common"
)

// CreateAccessGroup creates a new access group
func CreateAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		name             string
		organizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return
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
	if response.AffectedRows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

func UpdateAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result						sql.Result
		response					= SQLResponse{}
		tx								*sql.Tx
		name							string
		organizationName	string
		newName						string
		newOrganization		string
		updateStatements	[]string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}

	if newName, err = common.GetRequiredParam(params, "newName"); len(newName) > 0 {
		updateStatements = append(updateStatements, "name = " + newName)
	}
	if newOrganization, err = common.GetRequiredParam(params, "newOrganization"); len(newOrganization) > 0 {
		updateStatements = append(updateStatements, "organizationName = " + newOrganization)
	}

	if len(updateStatements) < 1 {
		return
	}

	if result, err = tx.Exec("UPDATE AccessGroup SET " + strings.Join(updateStatements, ", ") + " WHERE name = ? AND organizationName = ?;",
		name, organizationName); err != nil {
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

// DeleteAccessGroup deletes an organization
func DeleteAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
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
	if response.AffectedRows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// AddUserToAccessGroup adds a user to an existing organization
func AddUserToAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                      sql.Result
		response                    = SQLResponse{}
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
	if response.AffectedRows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// RemoveUserFromAccessGroup removes a user from an existing organization
func RemoveUserFromAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                      sql.Result
		response                    = SQLResponse{}
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
	if response.AffectedRows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// QueryAccessGroupOrganization queries from AccessGroup for organization
func QueryAccessGroupOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM AccessGroup "+
		"WHERE organizationName = ?;", organizationName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// QueryAccessGroupUserPairsOrganization queries user and access group pairings for organization
func QueryAccessGroupUserPairsOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM UserAccessGroupPairs "+
		"WHERE accessGroupOrganizationName = ?;", organizationName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}
