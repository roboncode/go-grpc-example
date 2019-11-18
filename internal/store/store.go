package store

type Store map[string]interface{}

func (s Store) Set(name string, value interface{}) {
	s[name] = value
}

func (s Store) Get(name string) interface{} {
	return s[name]
}

func NewStore() *Store {
	return &Store{}
}
