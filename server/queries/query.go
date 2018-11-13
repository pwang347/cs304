package queries

// SQLResponse is a JSON response for an SQL query
type SQLResponse struct {
	Rows int64 `json:"rows"`
}
