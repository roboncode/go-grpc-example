package service

import (
	"context"
	"example/generated"
	"github.com/golang/protobuf/ptypes/empty"
)

type httpService struct {
	example.UnimplementedHttpServiceServer
	PersonService example.PersonServiceServer
}

func (h httpService) CreatePerson(ctx context.Context, req *example.CreatePersonRequest) (*example.Person, error) {
	return h.PersonService.CreatePerson(ctx, req)
}

func (h httpService) GetPerson(ctx context.Context, req *example.GetPersonRequest) (*example.Person, error) {
	return h.PersonService.GetPerson(ctx, req)
}

func (h httpService) GetPersons(ctx context.Context, req *example.GetPersonsRequest) (*example.Persons, error) {
	return h.PersonService.GetPersons(ctx, req)
}

func (h httpService) UpdatePerson(ctx context.Context, req *example.UpdatePersonRequest) (*empty.Empty, error) {
	return h.PersonService.UpdatePerson(ctx, req)
}

func (h httpService) DeletePerson(ctx context.Context, req *example.DeletePersonRequest) (*empty.Empty, error) {
	return h.PersonService.DeletePerson(ctx, req)
}

func NewHttpService(personService example.PersonServiceServer) example.HttpServiceServer {
	return &httpService{
		PersonService: personService,
	}
}
