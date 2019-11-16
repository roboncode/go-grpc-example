package server

import (
	aaa "aaa/generated"
	"aaa/internal/store"
)

type Server struct {
	aaa.AppServer
	Store struct {
		Person *store.PersonStore
	}
}

func NewServer() *Server {
	return &Server{}
}
