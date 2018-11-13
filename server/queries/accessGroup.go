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
	if _, err = tx.Exec("INSERT INTO AccessGroup (name,organizationName) VALUES(?,?);",
		name, organizationName); err != nil {
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

// DeleteAccessGroup deletes an organization
func DeleteAccessGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
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
	if _, err = tx.Exec("DELETE FROM AccessGroup WHERE name=? AND organizationName=?;", name, organizationName); err != nil {
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
