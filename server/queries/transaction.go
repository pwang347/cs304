package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

// Create a new transaction
func CreateTransaction(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                sql.Result
		response              = SQLResponse{}
		tx                    *sql.Tx
		transactionNumStr     string
		transactionNum        int
		serviceSubTypeStr     string
		serviceSubType        int
		serviceSubServiceName string
		organizationName      string
		amountPaidStr         string
		amountPaid            int
		processedTimeStampStr string
		processedTimeStamp    int
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if transactionNumStr, err = common.GetRequiredParam(params, "transactionNumber"); err != nil {
		return
	}
	if transactionNum, err = strconv.Atoi(transactionNumStr); err != nil {
		return
	}
	if serviceSubTypeStr, err = common.GetRequiredParam(params, "serviceSubscriptionType"); err != nil {
		return
	}
	if serviceSubType, err = strconv.Atoi(serviceSubTypeStr); err != nil {
		return
	}
	if serviceSubServiceName, err = common.GetRequiredParam(params, "serviceSubscriptionServiceName"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if amountPaidStr, err = common.GetRequiredParam(params, "amountPaid"); err != nil {
		return
	}
	if amountPaid, err = strconv.Atoi(amountPaidStr); err != nil {
		return
	}
	if processedTimeStampStr, err = common.GetRequiredParam(params, "processedTimeStamp"); err != nil {
		return
	}
	if processedTimeStamp, err = strconv.Atoi(processedTimeStampStr); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO TRANSACTION (transactionNumber,serviceSubscriptionType,serviceSubscriptionServiceName, organizationName, amountPaid, processedTimestamp) VALUES(?,?,?,?,?,?);",
		transactionNum, serviceSubType, serviceSubServiceName, organizationName, amountPaid, processedTimeStamp); err != nil {
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

// Deletes a transaction
func DeleteTransaction(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result            sql.Result
		response          = SQLResponse{}
		tx                *sql.Tx
		transactionNumStr string
		transactionNum    int
		organizationName  string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if transactionNumStr, err = common.GetRequiredParam(params, "transactionNumber"); err != nil {
		return
	}
	if transactionNum, err = strconv.Atoi(transactionNumStr); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM TRANSACTION WHERE transactionNumber=? AND organizationName=?;",
		transactionNum, organizationName); err != nil {
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
