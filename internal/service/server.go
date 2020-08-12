package service

import (
	"example/generated"
	"example/internal/store"
)

type Server struct {
	example.AppServiceServer
	Store *store.Store
}

func NewServer(store *store.Store) *Server {
	return &Server{
		Store: store,
	}
}
