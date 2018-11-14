package queries

// SQLResponse is a JSON response for an SQL query
type SQLResponse struct {
	AffectedRows int64  `json:"affectedRows"`
	Data         string `json:"data"`
}
