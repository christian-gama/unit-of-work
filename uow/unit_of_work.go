package uow

// UnitOfWork is an interface that defines the functions necessary for a Unit of Work.
type UnitOfWork interface {
	Commit() error
	Rollback() error
	Transaction(func(uow UnitOfWork) error) error
}
