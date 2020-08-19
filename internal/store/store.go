package store

var _store *Store

type Store struct {
	Person PersonStore
}

func Instance() *Store {
	if _store == nil {
		_store = &Store{}
	}
	return _store
}
