package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

// CreateCreditCard creates a new creditCard
func CreateCreditCard(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response    = SQLResponse{Rows: 0}
		tx          *sql.Tx
		cardNumber  string
		cvc         string
		expiryDate  string // TODO: parse this
		cardTypeStr string
		cardType    int
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if cardNumber, err = common.GetRequiredParam(params, "cardNumber"); err != nil {
		return
	}
	if cvc, err = common.GetRequiredParam(params, "cvc"); err != nil {
		return
	}
	if expiryDate, err = common.GetRequiredParam(params, "expiryDate"); err != nil {
		return
	}
	if cardTypeStr, err = common.GetRequiredParam(params, "cardType"); err != nil {
		return
	}
	if cardType, err = strconv.Atoi(cardTypeStr); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO CreditCard (cardNumber,cvc,expiryDate,cardType) VALUES(?,?,?,?);",
		cardNumber, cvc, expiryDate, cardType); err != nil {
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

// DeleteCreditCard deletes a credit card
func DeleteCreditCard(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response   = SQLResponse{Rows: 0}
		tx         *sql.Tx
		cardNumber string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if cardNumber, err = common.GetRequiredParam(params, "cardNumber"); err != nil {
		return
	}
	if _, err = tx.Exec("DELETE FROM CreditCard WHERE cardNumber=?;", cardNumber); err != nil {
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
