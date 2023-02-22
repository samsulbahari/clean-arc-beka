package driven_repository

type PoolHandler interface {
	BeginTx() (Tx, error)
	Get(query string, args ...interface{}) Row
	Select(query string, args ...interface{}) (Rows, error)
	Exec(query string, args ...interface{}) error
}
