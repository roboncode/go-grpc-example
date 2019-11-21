package server

import (
	"aaa/internal/store"
	"aaa/pkg"
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
