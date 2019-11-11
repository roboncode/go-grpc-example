package server

import (
	aaa "aaa/generated"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) CreatePerson(ctx context.Context, req *aaa.Person) (*aaa.Person_Id, error) {
	return s.Store.Person.Create(context.Background(), req)
}

func (s *Server) GetPerson(ctx context.Context, req *aaa.Person_Id) (*aaa.Person, error) {
	return s.Store.Person.Get(context.Background(), req)
}

func (s *Server) GetPersons(ctx context.Context, req *aaa.Person_Filters) (*aaa.Persons, error) {
	return s.Store.Person.List(context.Background(), req)
}

func (s *Server) UpdatePerson(ctx context.Context, req *aaa.Person) (*aaa.Person, error) {
	return s.Store.Person.Update(context.Background(), req)
}

func (s *Server) DeletePerson(ctx context.Context, req *aaa.Person_Id) (*empty.Empty, error) {
	return s.Store.Person.Delete(context.Background(), req)
}

