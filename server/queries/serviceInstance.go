package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"

	"github.com/pwang347/cs304/server/common"
)

// CreateServiceInstance creates a new serviceInstance
func CreateServiceInstance(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		name             string
		regionName       string
		serviceName      string
		organizationName string
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
	if result, err = tx.Exec("INSERT INTO ServiceInstance (name, regionName, serviceName, organizationName) VALUES(?,?,?,?);",
		name, regionName, serviceName, organizationName); err != nil {
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

// DeleteServiceInstance deletes a serviceInstance
func DeleteServiceInstance(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		name             string
		serviceName      string
		organizationName string
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
	if result, err = tx.Exec("DELETE FROM ServiceInstance WHERE name=? AND serviceName=? AND organizationName=?;",
		name, serviceName, organizationName); err != nil {
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

// QueryServiceInstanceOrganization queries from serviceInstance for organization
func QueryServiceInstanceOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
		serviceName      string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM ServiceInstance "+
		"WHERE organizationName = ? AND serviceName = ?;", organizationName, serviceName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}
