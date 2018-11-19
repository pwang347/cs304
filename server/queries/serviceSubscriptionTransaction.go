package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

// CreateServiceSubscriptionTransaction creates a new service subscription
func CreateServiceSubscriptionTransaction(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result                sql.Result
		response              = SQLResponse{}
		tx                    *sql.Tx
		serviceTypeStr        string
		serviceType           int
		serviceName           string
		description           string
		organizationName      string
		activeUntil           string
		amountPaidStr         string
		amountPaid            int
		processedTimeStamp    string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if amountPaidStr, err = common.GetRequiredParam(params, "amountPaid"); err != nil {
		return
	}
	if amountPaid, err = strconv.Atoi(amountPaidStr); err != nil {
		return
	}
	if processedTimeStamp, err = common.GetRequiredParam(params, "processedTimestamp"); err != nil {
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
	if activeUntil, err = common.GetRequiredParam(params, "activeUntil"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO ServiceSubscriptionTransaction (amountPaid, processedTimestamp,type,serviceName,description, organizationName, activeUntil) VALUES(?,?,?,?,?,?,?);",
		amountPaid, processedTimeStamp, serviceType, serviceName, description, organizationName, activeUntil); err != nil {
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

// DeleteServiceSubscriptionTransactionByTransaction deletes a service subscription
func DeleteServiceSubscriptionTransactionByTransaction(db *sql.DB, params url.Values) (data []byte, err error) {
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

// ListAllActiveServiceSubscriptionTransactions lists all active service subscriptions that belongs to an organization
func ListAllActiveServiceSubscriptionTransactions(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM "+
		"ServiceSubscriptionTransaction SST, Service S, ServiceType ST, ServiceServiceTypePairs SSTP " +
		"WHERE S.name = SST.serviceName AND S.name = SSTP.serviceName " +
		"AND SSTP.serviceType = ST.type AND SST.type = SSTP.serviceType " +
		"AND organizationName = ? AND activeUntil >= NOW();",
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

// ListAllCompletedTransactions lists all completed transactions
func ListAllCompletedTransactions(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
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

// get all transactions for given month and year in an organization
func GetTransactionsForCurrentMonth(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
		currentMonth	 string
		currentYear	     string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if currentMonth, err = common.GetRequiredParam(params, "currentMonth"); err != nil {
		return
	}
	if currentYear, err = common.GetRequiredParam(params, "currentYear"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx,
		"SELECT * FROM ServiceSubscriptionTransaction WHERE organizationName = ? " +
		"AND MONTH(processedTimestamp) = ? AND YEAR(processedTimestamp) = ? ;",
		organizationName, currentMonth, currentYear); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// expired service subscriptions
func GetExpiredServiceSubscriptions(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx,
		"SELECT * FROM ServiceSubscriptionTransaction SST, ServiceServiceTypePairs SSTP, ServiceType ST"+
		" WHERE organizationName = ? " +
			"AND activeUntil < NOW() AND SST.serviceName = SSTP.serviceName AND" +
		" SST.type = SSTP.serviceType AND SSTP.serviceType = ST.type;",
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