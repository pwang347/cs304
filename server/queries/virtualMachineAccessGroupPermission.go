package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

func CreateVirtualMachineAccessGroupPermission(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result					sql.Result
		response				= SQLResponse{}
		tx							*sql.Tx
		vmIp						string
		orgName					string
		groupName				string
		accessLevelStr	string
		accessLevel			int
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if vmIp, err = common.GetRequiredParam(params, "vmIp"); err != nil {
		return
	}
	if orgName, err = common.GetRequiredParam(params, "orgName"); err != nil {
		return
	}
	if groupName, err = common.GetRequiredParam(params, "groupName"); err != nil {
		return
	}
	if accessLevelStr, err = common.GetRequiredParam(params, "accessLevel"); err != nil {
		return
	}
	if accessLevel, err = strconv.Atoi(accessLevelStr); err != nil {
		return
	}

	if result, err = tx.Exec("INSERT INTO VirtualMachineAccessGroupPermissions (VirtualMachineIpAddress,accessGroupOrganizationName,accessGroupName,accessLevel) VALUES (?,?,?,?);",
		vmIp, orgName, groupName, accessLevel); err != nil {
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

func DeleteVirtualMachineAccessGroupPermission(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result			sql.Result
		response		= SQLResponse{}
		tx					*sql.Tx
		vmIp				string
		orgName			string
		groupName		string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if vmIp, err = common.GetRequiredParam(params, "vmIp"); err != nil {
		return
	}
	if orgName, err = common.GetRequiredParam(params, "orgName"); err != nil {
		return
	}
	if groupName, err = common.GetRequiredParam(params, "groupName"); err != nil {
		return
	}

	if result, err = tx.Exec("DELETE FROM VirtualMachineAccessGroupPermissions WHERE VirtualMachineIpAddress = ? AND accessGroupOrganizationName = ? AND accessGroupName = ?;",
		vmIp, orgName, groupName); err != nil {
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

func QueryVirtualMachineAccessGroupPermissions(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result			sql.Result
		response		= SQLResponse{}
		tx					*sql.Tx
		vmIp				string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if vmIp, err = common.GetRequiredParam(params, "vmIp"); err != nil {
		return
	}

	if result, err = tx.Exec("SELECT * FROM VirtualMachineAccessGroupPermissions WHERE VirtualMachineIpAddress = ?;", vmIp); err != nil {
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
