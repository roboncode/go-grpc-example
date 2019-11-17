package store

type Store struct {
	Person *PersonStore
}

var singleton Store

func NewStore() *Store {
	return &singleton
}
