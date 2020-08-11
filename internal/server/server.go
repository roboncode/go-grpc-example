package server

import (
	"aaa/generated"
	"aaa/internal/store"
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
