package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"

	"github.com/pwang347/cs304/server/common"
)

// QueryAllBaseImages returns all baseImage rows
func QueryAllBaseImages(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response = SQLResponse{}
		tx       *sql.Tx
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM BaseImage;"); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}
