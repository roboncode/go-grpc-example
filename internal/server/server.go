package server

import (
	aaa "aaa/generated"
	"aaa/internal/stores"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	aaa.AppServer
	db *mongo.Database
	Stores struct{
		Person *stores.PersonStore
	}
}

func (s *Server) InitMongoStore(db *mongo.Database) {
	s.Stores.Person = stores.NewPersonStore(db)
}

func NewServer() *Server {
	return &Server{}
}
