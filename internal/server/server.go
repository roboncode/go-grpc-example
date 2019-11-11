package server

import (
	aaa "aaa/generated"
	"aaa/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	aaa.AppServer
	db    *mongo.Database
	Store struct{
		Person *store.PersonStore
	}
}

func (s *Server) InitMongoStore(db *mongo.Database) {
	s.Store.Person = store.NewPersonStore(db)
}

func NewServer() *Server {
	return &Server{}
}
