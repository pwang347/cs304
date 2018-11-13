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
		response                        = SQLResponse{Rows: 0}
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
	if _, err = tx.Exec("INSERT INTO ServiceInstanceConfiguration (configKey,serviceInstanceName,serviceInstanceServiceName,"+
		"serviceInstanceOrganizationName,data) VALUES(?,?,?,?,?);",
		configKey, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName, dataStr); err != nil {
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

// DeleteServiceInstanceConfiguration deletes a service instance configuration
func DeleteServiceInstanceConfiguration(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response                        = SQLResponse{Rows: 0}
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
	if _, err = tx.Exec("DELETE FROM ServiceInstanceConfiguration WHERE configKey=? AND serviceInstanceName=? AND serviceInstanceServiceName=?"+
		" AND serviceInstanceOrganizationName=?;", configKey, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName); err != nil {
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
