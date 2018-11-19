package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"

	"github.com/pwang347/cs304/server/common"
)

// CreateServiceInstanceKey creates a new service instance configuration
func CreateServiceInstanceKey(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                          sql.Result
		response                        = SQLResponse{}
		tx                              *sql.Tx
		keyValue                        = common.GenerateRandomHash(16)
		serviceInstanceName             string
		serviceInstanceServiceName      string
		serviceInstanceOrganizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if serviceInstanceName, err = common.GetRequiredParam(params, "serviceInstanceName"); err != nil {
		return
	}
	if serviceInstanceServiceName, err = common.GetRequiredParam(params, "serviceInstanceServiceName"); err != nil {
		return
	}
	if serviceInstanceOrganizationName, err = common.GetRequiredParam(params, "serviceInstanceOrganizationName"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO ServiceInstanceKey (keyValue,activeUntil,serviceInstanceName,serviceInstanceServiceName,serviceInstanceOrganizationName) "+
		"VALUES(?,(NOW() + INTERVAL 1 DAY),?,?,?);",
		keyValue, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName); err != nil {
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

// DeleteServiceInstanceKey deletes a service instance configuration
func DeleteServiceInstanceKey(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                          sql.Result
		response                        = SQLResponse{}
		tx                              *sql.Tx
		keyValue                        string
		serviceInstanceName             string
		serviceInstanceServiceName      string
		serviceInstanceOrganizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if keyValue, err = common.GetRequiredParam(params, "keyValue"); err != nil {
		return
	}
	if serviceInstanceName, err = common.GetRequiredParam(params, "serviceInstanceName"); err != nil {
		return
	}
	if serviceInstanceServiceName, err = common.GetRequiredParam(params, "serviceInstanceServiceName"); err != nil {
		return
	}
	if serviceInstanceOrganizationName, err = common.GetRequiredParam(params, "serviceInstanceOrganizationName"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM ServiceInstanceKey "+
		"WHERE keyValue=? AND serviceInstanceName=? AND serviceInstanceServiceName=? AND serviceInstanceOrganizationName=?;",
		keyValue, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName); err != nil {
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

// UpdateServiceInstanceKey updates the details of a service instance configuration
func UpdateServiceInstanceKey(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                          sql.Result
		response                        = SQLResponse{}
		tx                              *sql.Tx
		keyValue                        string
		serviceInstanceName             string
		serviceInstanceServiceName      string
		serviceInstanceOrganizationName string
		newKeyValue                     string
		newActiveUntil                  string // TODO
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if keyValue, err = common.GetRequiredParam(params, "keyValue"); err != nil {
		return
	}
	if serviceInstanceName, err = common.GetRequiredParam(params, "serviceInstanceName"); err != nil {
		return
	}
	if serviceInstanceServiceName, err = common.GetRequiredParam(params, "serviceInstanceServiceName"); err != nil {
		return
	}
	if serviceInstanceOrganizationName, err = common.GetRequiredParam(params, "serviceInstanceOrganizationName"); err != nil {
		return
	}
	if newKeyValue, err = common.GetRequiredParam(params, "newKeyValue"); err != nil {
		return
	}
	if newActiveUntil, err = common.GetRequiredParam(params, "newActiveUntil"); err != nil {
		return
	}
	if result, err = tx.Exec("UPDATE ServiceInstanceKey "+
		"SET keyValue=?, activeUntil=? "+
		"WHERE keyValue=? AND serviceInstanceName=? AND serviceInstanceServiceName=? AND serviceInstanceOrganizationName=?;",
		newKeyValue, newActiveUntil, keyValue, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName); err != nil {
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

// QueryServiceInstanceKeys queries all keys for a service instance
func QueryServiceInstanceKeys(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response                        = SQLResponse{}
		tx                              *sql.Tx
		serviceInstanceName             string
		serviceInstanceServiceName      string
		serviceInstanceOrganizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if serviceInstanceName, err = common.GetRequiredParam(params, "serviceInstanceName"); err != nil {
		return
	}
	if serviceInstanceServiceName, err = common.GetRequiredParam(params, "serviceInstanceServiceName"); err != nil {
		return
	}
	if serviceInstanceOrganizationName, err = common.GetRequiredParam(params, "serviceInstanceOrganizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM ServiceInstanceKey "+
		"WHERE serviceInstanceName = ? AND serviceInstanceServiceName = ? AND serviceInstanceOrganizationName = ?;",
		serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}
