package services

import (
	"context"
	"example/generated"
	"example/internal/models"
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

func NewPersonService() example.PersonServiceServer {
	return &personService{}
}

func (p *personService) CreatePerson(_ context.Context, req *example.CreatePersonRequest) (*example.Person, error) {
	var personStore = store.Instance().Person
	var person = models.Person{
		Status: int32(req.Status),
		Name:   req.Name,
		Email:  req.Email,
	}
	err := personStore.CreatePerson(&person)
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
	var personStore = store.Instance().Person
	person, err := personStore.GetPerson(req.Id)
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
	var personStore = store.Instance().Person
	var filters = store.PersonFilters{
		Status: int32(req.Status),
		Skip:   req.Skip,
		Limit:  req.Limit,
	}

	persons, err := personStore.GetPersons(&filters)
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
	var personStore = store.Instance().Person
	var person = models.Person{
		Status: int32(req.Status),
		Name:   req.Name,
		Email:  req.Email,
	}
	var err = personStore.UpdatePerson(req.Id, &person)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error updated person with id {%s}: %v", req.Id, err)
	}
	return &empty.Empty{}, nil
}

func (p *personService) DeletePerson(_ context.Context, req *example.DeletePersonRequest) (*empty.Empty, error) {
	var personStore = store.Instance().Person
	var err = personStore.DeleteRequest(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error deleting person with id {%s}: %v", req.Id, err)
	}
	return &empty.Empty{}, nil
}
