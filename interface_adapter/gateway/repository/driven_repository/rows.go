package driven_repository

type Rows interface {
	Close()

	Next() bool

	Scan(dest ...any) error
}
