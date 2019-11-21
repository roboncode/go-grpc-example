package server

import (
	"aaa/internal/store"
	"aaa/pkg"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) PersonStore() *store.PersonStore {
	return s.Store.Get(store.PersonStoreName).(*store.PersonStore)
}

func (s *Server) CreatePerson(ctx context.Context, req *pkg.Person) (*pkg.Person_Id, error) {
	return s.PersonStore().Create(ctx, req)
}

func (s *Server) GetPerson(ctx context.Context, req *pkg.Person_Id) (*pkg.Person, error) {
	return s.PersonStore().Get(ctx, req)
}

func (s *Server) GetPersons(ctx context.Context, req *pkg.Person_Filters) (*pkg.Persons, error) {
	return s.PersonStore().List(ctx, req)
}

func (s *Server) UpdatePerson(ctx context.Context, req *pkg.Person) (*pkg.Person, error) {
	return s.PersonStore().Update(ctx, req)
}

func (s *Server) DeletePerson(ctx context.Context, req *pkg.Person_Id) (*empty.Empty, error) {
	return s.PersonStore().Delete(ctx, req)
}
