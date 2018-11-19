package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

// CreateVirtualMachine creates a new virtual machine
func CreateVirtualMachine(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                    sql.Result
		response                  = SQLResponse{}
		tx                        *sql.Tx
		ipAddress                 = common.GenerateIPAddress()
		description               string
		coresStr                  string
		cores                     int
		diskSpaceStr              string
		diskSpace                 int
		ramStr                    string
		ram                       int
		baseImageOS               string
		baseImageVersion          string
		regionName                string
		organizationName          string
		virtualMachineServiceName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if description, err = common.GetRequiredParam(params, "description"); err != nil {
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
		"INSERT INTO VirtualMachine "+
			"(description, ipAddress, state, cores, diskSpace, ram, baseImageOs, baseImageVersion, regionName, organizationName, virtualMachineServiceName)"+
			" VALUES(?,?,?,?,?,?,?,?,?,?,?);",
		description, ipAddress, 0, cores, diskSpace, ram, baseImageOS, baseImageVersion, regionName, organizationName, virtualMachineServiceName); err != nil {
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

// DeleteVirtualMachine deletes a virtual machine
func DeleteVirtualMachine(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result    sql.Result
		response  = SQLResponse{}
		tx        *sql.Tx
		ipAddress string
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

// UpdateVirtualMachineState updates a virtual machine state
func UpdateVirtualMachineState(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result    sql.Result
		response  = SQLResponse{}
		tx        *sql.Tx
		ipAddress string
		stateStr  string
		state     int
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}

	if ipAddress, err = common.GetRequiredParam(params, "ipAddress"); err != nil {
		return
	}
	if stateStr, err = common.GetRequiredParam(params, "state"); err != nil {
		return
	}
	if state, err = strconv.Atoi(stateStr); err != nil {
		return
	}
	if result, err = tx.Exec("UPDATE VirtualMachine SET state=? WHERE ipAddress = ?;",
		state, ipAddress); err != nil {
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

// UpdateVirtualMachine updates a virtual machine
func UpdateVirtualMachine(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		ipAddress        string
		description      string
		coresStr         string
		cores            int
		diskSpaceStr     string
		diskSpace        int
		ramStr           string
		ram              int
		baseImageOS      string
		baseImageVersion string
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

	if result, err = tx.Exec("UPDATE VirtualMachine SET description=?, cores=?, diskSpace=?, ram=?, baseImageOs=?, baseImageVersion=? "+
		"WHERE ipAddress = ?;",
		description, cores, diskSpace, ram, baseImageOS, baseImageVersion, ipAddress); err != nil {
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

// QueryVirtualMachineOrganization queries from serviceInstance for organization
func QueryVirtualMachineOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
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
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM VirtualMachine "+
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

// QueryVirtualMachineServiceOrganization queries from virtual machine for organization and service pair
func QueryVirtualMachineServiceOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response                  = SQLResponse{}
		tx                        *sql.Tx
		organizationName          string
		virtualMachineServiceName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if virtualMachineServiceName, err = common.GetRequiredParam(params, "virtualMachineServiceName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM VirtualMachine "+
		"WHERE organizationName = ? AND virtualMachineServiceName = ?;", organizationName, virtualMachineServiceName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

func QueryMostPopularVMImage(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response = SQLResponse{}
		tx       *sql.Tx
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}

	if response.Data, response.AffectedRows, err = common.QueryJSON(tx,
		"SELECT baseImageOs AS os, baseImageVersion AS version FROM VirtualMachine"+
			"GROUP BY baseImageOs, baseImageVersion HAVING COUNT(*) = ("+
			"SELECT MAX(osCount) FROM ("+
			"SELECT COUNT(*) AS osCount FROM VirtualMachine GROUP BY baseImageOs, baseImageVersion))"); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}
