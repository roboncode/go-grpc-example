package service

import (
	"context"
	"example/generated"
	"example/internal/store"
	"example/internal/store/mocks"
	"example/util/transform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var now = time.Now()
var personItem = store.Person{
	Id:        transform.ToObjectId("5f35832b7913cffd5a329af7"),
	Status:    int32(example.Status_ACTIVE),
	Name:      "John Smith",
	Email:     "john.smith@fakemail.com",
	CreatedAt: &now,
	UpdatedAt: &now,
}

func newPersonStore() store.PersonStore {
	var mockStore = mocks.PersonStore{}
	//CreatePerson(person *Person) error
	mockStore.On("CreatePerson", mock.Anything).Return(func(person *store.Person) error {
		person.Id = personItem.Id
		person.CreatedAt = personItem.CreatedAt
		person.UpdatedAt = personItem.UpdatedAt
		return nil
	})
	//GetPerson(id string) (*Person, error)
	mockStore.On("GetPerson", mock.Anything).Return(&personItem, nil)
	//GetPersons(filters *PersonFilters) ([]Person, error)
	mockStore.On("GetPersons", mock.Anything).Return(func(filters *store.PersonFilters) []store.Person {
		var persons = make([]store.Person, 0)
		persons = append(persons, personItem)
		return persons
	}, nil)
	//UpdatePerson(id string, person *Person) error
	mockStore.On("UpdatePerson", mock.Anything, mock.Anything).Return(nil)
	//DeleteRequest(id string) error
	mockStore.On("DeleteRequest", mock.Anything).Return(nil)
	return &mockStore
}

type TestServiceSuite struct {
	suite.Suite
	server example.PersonServiceServer
}

func (p *TestServiceSuite) SetupTest() {
	var personStore = store.NewStore()
	personStore.Set(store.PersonStoreName, newPersonStore())
	p.server = NewPersonService(personStore)
}

func (p *TestServiceSuite) TestCreatePerson() {
	request := example.CreatePersonRequest{
		Status: 0,
		Name:   "John Smith",
		Email:  "john.smith@fakemail.com",
	}
	response, err := p.server.CreatePerson(context.Background(), &request)
	assert.Nil(p.T(), err)
	assert.NotNil(p.T(), response)
}

func (p *TestServiceSuite) TestGetPersons() {
	request := example.GetPersonsRequest{}
	response, err := p.server.GetPersons(context.Background(), &request)
	assert.Nil(p.T(), err)
	assert.NotNil(p.T(), response)
}

func (p *TestServiceSuite) TestUpdatePerson() {
	request := example.UpdatePersonRequest{
		Id:     personItem.Id.Hex(),
		Status: example.Status(personItem.Status),
		Name:   personItem.Name,
		Email:  personItem.Email,
	}
	response, err := p.server.UpdatePerson(context.Background(), &request)
	assert.Nil(p.T(), err)
	assert.NotNil(p.T(), response)
}

func (p *TestServiceSuite) TestDeletePerson() {
	request := example.DeletePersonRequest{
		Id: personItem.Id.Hex(),
	}
	response, err := p.server.DeletePerson(context.Background(), &request)
	assert.Nil(p.T(), err)
	assert.NotNil(p.T(), response)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestServiceSuite))
}
