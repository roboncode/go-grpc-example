package server

import (
	aaa "aaa/generated"
	"aaa/internal/store"
)

type Server struct {
	aaa.AppServer
	Store *store.Store
}

func NewServer(store *store.Store) *Server {
	return &Server{
		Store: store,
	}
}
