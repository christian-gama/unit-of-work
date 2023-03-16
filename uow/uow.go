package uow

type UnitOfWork interface {
	Commit() error
	Rollback() error
	Begin() error
}
