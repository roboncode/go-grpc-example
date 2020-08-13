package service

import (
	"context"
	"example/generated"
	"example/internal/store"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func ToTimestamp(t time.Time) *timestamp.Timestamp {
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil
	}
	return ts
}

func (p *Server) PersonStore() store.PersonStore {
	return p.Store.Get(store.PersonStoreName).(store.PersonStore)
}

func (p *Server) CreatePerson(_ context.Context, req *example.CreatePersonRequest) (*example.Person, error) {
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
		CreatedAt: ToTimestamp(*person.CreatedAt),
		UpdatedAt: ToTimestamp(*person.UpdatedAt),
	}, nil
}

func (p *Server) GetPerson(_ context.Context, req *example.GetPersonRequest) (*example.Person, error) {
	person, err := p.PersonStore().GetPerson(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not find Person with Object Id %s: %v", req.Id, err)
	}
	return &example.Person{
		Id:        person.Id.Hex(),
		Status:    example.Status(person.Status),
		Name:      person.Name,
		Email:     person.Email,
		CreatedAt: ToTimestamp(*person.CreatedAt),
		UpdatedAt: ToTimestamp(*person.UpdatedAt),
	}, nil
}

func (p *Server) GetPersons(_ context.Context, req *example.GetPersonsRequest) (*example.Persons, error) {
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
			CreatedAt: ToTimestamp(*person.CreatedAt),
			UpdatedAt: ToTimestamp(*person.UpdatedAt),
		}
		response.Items = append(response.Items, &item)
	}
	return &response, nil
}

func (p *Server) UpdatePerson(_ context.Context, req *example.UpdatePersonRequest) (*empty.Empty, error) {
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

func (p *Server) DeletePerson(_ context.Context, req *example.DeletePersonRequest) (*empty.Empty, error) {
	var err = p.PersonStore().DeleteRequest(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error deleting person with id {%s}: %v", req.Id, err)
	}
	return &empty.Empty{}, nil
}
