package driven_repository

type Row interface {
	Scan(dest ...any) error
}
