package queries

import (
	"database/sql"
	"encoding/json"
	"github.com/pwang347/cs304/server/common"
	"net/url"
	"strconv"
)

// CreateServiceSubscriptionTransaction creates a new service subscription
func CreateServiceSubscriptionTransaction(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		serviceTypeStr   string
		serviceType		 int
		serviceName		 string
		description		 string
		organizationName string
		activeUntilStr	 string
		activeUntil		 int
		amountPaidStr			string
		amountPaid		 		int
		processedTimeStampStr	string
		processedTimeStamp		int
		transactionNumStr		string
		transactionNum			int
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
	if serviceTypeStr, err = common.GetRequiredParam(params, "serviceType"); err != nil {
		return
	}
	if serviceType, err = strconv.Atoi(serviceTypeStr); err != nil {
		return
	}
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if description, err = common.GetRequiredParam(params, "description"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if activeUntilStr, err = common.GetRequiredParam(params, "activeUntil"); err != nil {
		return
	}
	if activeUntil, err = strconv.Atoi(activeUntilStr); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO ServiceSubscriptionTransaction (transactionNumber, amountPaid, processedTimestamp, type,serviceName,description, organizationName, activeUntil) VALUES(?,?,?,?,?,?,?,?);",
		transactionNum, amountPaid, processedTimeStamp, serviceType, serviceName, description, organizationName, activeUntil); err != nil {
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

// deletes a service subscription
func DeleteServiceSubscriptionTransactionByTransaction(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           		sql.Result
		response         		= SQLResponse{}
		tx               		*sql.Tx
		transactionNumStr		string
		transactionNum			int
		organizationName		string
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
	if result, err = tx.Exec("DELETE FROM ServiceSubscriptionTransaction WHERE transactionNumber=? AND organizationName=?;",
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

// lists all active service subscriptions that belongs to an organization
func ListAllActiveServiceSubscriptionTransactions(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         		= SQLResponse{}
		tx               		*sql.Tx
		organizationName		string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM " +
		"ServiceSubscriptionTransaction WHERE organizationName = ? AND activeUntil >= NOW();",
		organizationName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// lists all completed transactions
func ListAllCompletedTransactions(db *sql.DB, params url.Values) (data []byte, err error){
	var (
		response         		= SQLResponse{}
		tx               		*sql.Tx
		organizationName		string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM ServiceSubscriptionTransaction WHERE organizationName = ?;",
		organizationName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}