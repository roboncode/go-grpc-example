package service

import (
	"context"
	"example/generated"
	"example/internal/store"
	"example/util/transform"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type personService struct {
	example.UnimplementedPersonServiceServer
	Store store.Store
}

func NewPersonService(store store.Store) example.PersonServiceServer {
	return &personService{
		Store: store,
	}
}

func (p *personService) PersonStore() store.PersonStore {
	return p.Store.Get(store.PersonStoreName).(store.PersonStore)
}

func (p *personService) CreatePerson(_ context.Context, req *example.CreatePersonRequest) (*example.Person, error) {
	var person = store.Person{
		Status: int32(req.Status),
		Name:   req.Name,
		Email:  req.Email,
	}
	err := p.PersonStore().CreatePerson(&person)
	if err != nil {
		return nil, err
	}

	return &example.Person{
		Id:        person.Id.Hex(),
		Status:    example.Status(person.Status),
		Name:      person.Name,
		Email:     person.Email,
		CreatedAt: transform.ToTimestamp(*person.CreatedAt),
		UpdatedAt: transform.ToTimestamp(*person.UpdatedAt),
	}, nil
}

func (p *personService) GetPerson(_ context.Context, req *example.GetPersonRequest) (*example.Person, error) {
	person, err := p.PersonStore().GetPerson(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not find Person with Object Id %s: %v", req.Id, err)
	}
	return &example.Person{
		Id:        person.Id.Hex(),
		Status:    example.Status(person.Status),
		Name:      person.Name,
		Email:     person.Email,
		CreatedAt: transform.ToTimestamp(*person.CreatedAt),
		UpdatedAt: transform.ToTimestamp(*person.UpdatedAt),
	}, nil
}

func (p *personService) GetPersons(_ context.Context, req *example.GetPersonsRequest) (*example.Persons, error) {
	var filters = store.PersonFilters{
		Status: int32(req.Status),
		Skip:   req.Skip,
		Limit:  req.Limit,
	}

	persons, err := p.PersonStore().GetPersons(&filters)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error retrieving persons: %s", err)
	}

	var response example.Persons
	for _, person := range persons {
		var item = example.Person{
			Id:        person.Id.Hex(),
			Status:    example.Status(person.Status),
			Name:      person.Name,
			Email:     person.Email,
			CreatedAt: transform.ToTimestamp(*person.CreatedAt),
			UpdatedAt: transform.ToTimestamp(*person.UpdatedAt),
		}
		response.Items = append(response.Items, &item)
	}
	return &response, nil
}

func (p *personService) UpdatePerson(_ context.Context, req *example.UpdatePersonRequest) (*empty.Empty, error) {
	var person = store.Person{
		Status: int32(req.Status),
		Name:   req.Name,
		Email:  req.Email,
	}
	var err = p.PersonStore().UpdatePerson(req.Id, &person)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error updated person with id {%s}: %v", req.Id, err)
	}
	return &empty.Empty{}, nil
}

func (p *personService) DeletePerson(_ context.Context, req *example.DeletePersonRequest) (*empty.Empty, error) {
	var err = p.PersonStore().DeleteRequest(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error deleting person with id {%s}: %v", req.Id, err)
	}
	return &empty.Empty{}, nil
}
