// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import emptypb "google.golang.org/protobuf/types/known/emptypb"
import example "example/generated"
import mock "github.com/stretchr/testify/mock"

// HttpServiceServer is an autogenerated mock type for the HttpServiceServer type
type HttpServiceServer struct {
	mock.Mock
}

// CreatePerson provides a mock function with given fields: _a0, _a1
func (_m *HttpServiceServer) CreatePerson(_a0 context.Context, _a1 *example.CreatePersonRequest) (*example.Person, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *example.Person
	if rf, ok := ret.Get(0).(func(context.Context, *example.CreatePersonRequest) *example.Person); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*example.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *example.CreatePersonRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePerson provides a mock function with given fields: _a0, _a1
func (_m *HttpServiceServer) DeletePerson(_a0 context.Context, _a1 *example.DeletePersonRequest) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *example.DeletePersonRequest) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *example.DeletePersonRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPerson provides a mock function with given fields: _a0, _a1
func (_m *HttpServiceServer) GetPerson(_a0 context.Context, _a1 *example.GetPersonRequest) (*example.Person, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *example.Person
	if rf, ok := ret.Get(0).(func(context.Context, *example.GetPersonRequest) *example.Person); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*example.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *example.GetPersonRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPersons provides a mock function with given fields: _a0, _a1
func (_m *HttpServiceServer) GetPersons(_a0 context.Context, _a1 *example.GetPersonsRequest) (*example.Persons, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *example.Persons
	if rf, ok := ret.Get(0).(func(context.Context, *example.GetPersonsRequest) *example.Persons); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*example.Persons)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *example.GetPersonsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePerson provides a mock function with given fields: _a0, _a1
func (_m *HttpServiceServer) UpdatePerson(_a0 context.Context, _a1 *example.UpdatePersonRequest) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *example.UpdatePersonRequest) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *example.UpdatePersonRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
