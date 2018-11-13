package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

// CreateServiceInstance creates a new serviceInstance
func CreateServiceInstance(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response             = SQLResponse{Rows: 0}
		tx                   *sql.Tx
		name                string
		regionName          string
		serviceName         string
		organizationName    string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if regionName, err = common.GetRequiredParam(params, "regionName"); err != nil {
		return
	}
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO ServiceInstance (name, regionName, serviceName, organizationName) VALUES(?,?,?,?);",
		name, regionName, serviceName, organizationName); err != nil {
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

// DeleteServiceInstance deletes a serviceInstance
func DeleteServiceInstance(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response = SQLResponse{Rows: 0}
		tx       *sql.Tx
		name                string
		serviceName         string
		organizationName    string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if _, err = tx.Exec("DELETE FROM ServiceInstance WHERE name=? AND serviceName=? AND organizationName=? AND;", emailAddress); err != nil {
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
