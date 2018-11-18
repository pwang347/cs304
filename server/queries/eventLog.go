package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"time"

	"github.com/pwang347/cs304/server/common"
)

func QueryEventLogsForVirtualMachine(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response = SQLResponse{}
		tx       *sql.Tx
		vmIp     string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if vmIp, err = common.GetRequiredParam(params, "vmIp"); err != nil {
		return
	}

	if response.Data, response.AffectedRows, err = common.QueryJSON(tx,
		"SELECT * FROM EventLog WHERE VirtualMachineIpAddress = ?", vmIp); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

func CreateEventLog(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result    sql.Result
		response  = SQLResponse{}
		tx        *sql.Tx
		logData   string
		eventType string
		vmIp      string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if logData, err = common.GetRequiredParam(params, "logData"); err != nil {
		return
	}
	if eventType, err = common.GetRequiredParam(params, "eventType"); err != nil {
		return
	}
	if eventType, err = common.GetRequiredParam(params, "vmIp"); err != nil {
		return
	}

	t := time.Now()
	ts := t.Format("2006-01-02 15:04:05")

	if result, err = tx.Exec("INSERT INTO EventLog (logNumber, timestamp, data, eventType, VirtualMachineIpAddress) "+
		"VALUES((SELECT COALESCE(MAX(logNumber)+1,1) FROM EventLog WHERE VirtualMachineIpAddress = "+vmIp+
		"),?,?,?,?);",
		ts, logData, eventType, vmIp); err != nil {
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
