package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"

	"github.com/pwang347/cs304/server/common"
)

// CreateServiceInstanceConfiguration creates a new service instance configuration
func CreateServiceInstanceConfiguration(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                          sql.Result
		response                        = SQLResponse{}
		tx                              *sql.Tx
		configKey                       string
		serviceInstanceName             string
		serviceInstanceServiceName      string
		serviceInstanceOrganizationName string
		dataStr                         string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if configKey, err = common.GetRequiredParam(params, "configKey"); err != nil {
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
	if dataStr, err = common.GetRequiredParam(params, "data"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO ServiceInstanceConfiguration (configKey,serviceInstanceName,serviceInstanceServiceName,serviceInstanceOrganizationName,data) "+
		"VALUES(?,?,?,?,?);",
		configKey, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName, dataStr); err != nil {
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

// DeleteServiceInstanceConfiguration deletes a service instance configuration
func DeleteServiceInstanceConfiguration(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                          sql.Result
		response                        = SQLResponse{}
		tx                              *sql.Tx
		configKey                       string
		serviceInstanceName             string
		serviceInstanceServiceName      string
		serviceInstanceOrganizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if configKey, err = common.GetRequiredParam(params, "configKey"); err != nil {
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
	if result, err = tx.Exec("DELETE FROM ServiceInstanceConfiguration "+
		"WHERE configKey=? AND serviceInstanceName=? AND serviceInstanceServiceName=? AND serviceInstanceOrganizationName=?;",
		configKey, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName); err != nil {
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

// UpdateServiceInstanceConfiguration updates the details of a service instance configuration
func UpdateServiceInstanceConfiguration(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                          sql.Result
		response                        = SQLResponse{}
		tx                              *sql.Tx
		configKey                       string
		serviceInstanceName             string
		serviceInstanceServiceName      string
		serviceInstanceOrganizationName string
		newConfigKey                    string
		newData                         string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if configKey, err = common.GetRequiredParam(params, "configKey"); err != nil {
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
	if newConfigKey, err = common.GetRequiredParam(params, "newConfigKey"); err != nil {
		return
	}
	if newData, err = common.GetRequiredParam(params, "newData"); err != nil {
		return
	}
	if result, err = tx.Exec("UPDATE ServiceInstanceConfiguration "+
		"SET configKey=?, data=? "+
		"WHERE configKey=? AND serviceInstanceName=? AND serviceInstanceServiceName=? AND serviceInstanceOrganizationName=?;",
		newConfigKey, newData, configKey, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName); err != nil {
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
