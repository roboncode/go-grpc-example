package server

import (
	"context"
	"example/generated"
	"example/internal/store"
	"github.com/golang/protobuf/ptypes/empty"
)

func (p *Server) PersonStore() *store.PersonStore {
	return p.Store.Get(store.PersonStoreName).(*store.PersonStore)
}

func (p *Server) CreatePerson(ctx context.Context, req *example.CreatePersonRequest) (*example.Person, error) {
	return p.PersonStore().CreatePerson(ctx, req)
}

func (p *Server) GetPerson(ctx context.Context, req *example.GetPersonRequest) (*example.Person, error) {
	return p.PersonStore().GetPerson(ctx, req)
}

func (p *Server) GetPersons(ctx context.Context, req *example.GetPersonsRequest) (*example.Persons, error) {
	return p.PersonStore().GetPersons(ctx, req)
}

func (p *Server) UpdatePerson(ctx context.Context, req *example.UpdatePersonRequest) (*empty.Empty, error) {
	return p.PersonStore().UpdatePerson(ctx, req)
}

func (p *Server) DeletePerson(ctx context.Context, req *example.DeletePersonRequest) (*empty.Empty, error) {
	return p.PersonStore().DeleteRequest(ctx, req)
}
