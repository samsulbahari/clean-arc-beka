package driven_repository

type Tx interface {
	Commit() error
	Rollback() error
	Get(query string, args ...interface{}) Row
	Select(query string, args ...interface{}) (Rows, error)
	Exec(query string, args ...interface{}) error
}
