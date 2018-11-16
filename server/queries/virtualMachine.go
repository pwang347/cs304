package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

// creates a new virtual machine
func CreateVirtualMachine(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		ipAddress        string
		description      string
		stateStr         string
		state			 int
		coresStr		 string
		cores			 int
		diskSpaceStr	 string
		diskSpace		 int
		ramStr			 string
		ram				 int
		baseImageOS		 string
		baseImageVersion string
		regionName		 string
		organizationName string
		virtualMachineServiceName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if ipAddress, err = common.GetRequiredParam(params, "ipAddress"); err != nil {
		return
	}
	if description, err = common.GetRequiredParam(params, "description"); err != nil {
		return
	}
	if stateStr, err = common.GetRequiredParam(params, "state"); err != nil {
		return
	}
	if state, err = strconv.Atoi(stateStr); err != nil {
		return
	}
	if coresStr, err = common.GetRequiredParam(params, "cores"); err != nil {
		return
	}
	if cores, err = strconv.Atoi(coresStr); err != nil {
		return
	}
	if diskSpaceStr, err = common.GetRequiredParam(params, "diskSpace"); err != nil {
		return
	}
	if diskSpace, err = strconv.Atoi(diskSpaceStr); err != nil {
		return
	}
	if ramStr, err = common.GetRequiredParam(params, "ram"); err != nil {
		return
	}
	if ram, err = strconv.Atoi(ramStr); err != nil {
		return
	}
	if baseImageOS, err = common.GetRequiredParam(params, "baseImageOs"); err != nil {
		return
	}
	if baseImageVersion, err = common.GetRequiredParam(params, "baseImageVersion"); err != nil {
		return
	}
	if regionName, err = common.GetRequiredParam(params, "regionName"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if virtualMachineServiceName, err = common.GetRequiredParam(params, "virtualMachineServiceName"); err != nil {
		return
	}
	if result, err = tx.Exec(
		"INSERT INTO VirtualMachine " +
			"(description, ipAddress, state, cores, diskSpace, ram, baseImageOs, baseImageVersion, regionName, organizationName, virtualMachineServiceName)" +
			" VALUES(?,?,?,?,?,?,?,?,?,?,?);",
		description, ipAddress, state, cores, diskSpace, ram, baseImageOS, baseImageVersion, regionName, organizationName, virtualMachineServiceName); err != nil {
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

// deletes a virtual machine
func DeleteVirtualMachine(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		ipAddress		 string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if ipAddress, err = common.GetRequiredParam(params, "ipAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM VirtualMachine WHERE ipAddress=?;",
		ipAddress); err != nil {
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
