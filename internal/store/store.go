package store

type Store interface {
	Set(name string, value interface{})
	Get(name string) interface{}
}

type store map[string]interface{}

func (s store) Set(name string, value interface{}) {
	s[name] = value
}

func (s store) Get(name string) interface{} {
	return s[name]
}

func NewStore() Store {
	return &store{}
}
