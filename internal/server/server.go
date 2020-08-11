package server

import (
	"example/generated"
	"example/internal/store"
)

type Server struct {
	pkg.AppServer
	Store *store.Store
}

func NewServer(store *store.Store) *Server {
	return &Server{
		Store: store,
	}
}
