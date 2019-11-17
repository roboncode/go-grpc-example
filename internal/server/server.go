package server

import (
	aaa "aaa/generated"
	"aaa/internal/store"
)

type Server struct {
	aaa.AppServer
	Store *store.Store
}

func NewServer() *Server {
	return &Server{
		Store: store.NewStore(),
	}
}
