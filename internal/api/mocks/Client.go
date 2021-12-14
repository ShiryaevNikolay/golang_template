package mocks

import (
	api "example/internal/api"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerates mock type for the Client type
type Client struct {
	mock.Mock
}

// GetJoke provides a mock funcrion with given fields:
func (_m *Client) GetJoke() (*api.JokeResponse, error) {
	ret := _m.Called()

	var r0 *api.JokeResponse
	if rf, ok := ret.Get(0).(func() *api.JokeResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.JokeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
