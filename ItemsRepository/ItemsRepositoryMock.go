package itemsrepository

import (
	mock "github.com/stretchr/testify/mock"
)

func NewRepositoryMock() *itemsRepositoryMock {
	return &itemsRepositoryMock{}
}

type itemsRepositoryMock struct {
	mock.Mock
}

func (_m *itemsRepositoryMock) GetItems() ([]Output, error) {
	ret := _m.Called()

	var r0 []Output
	if rf, ok := ret.Get(0).(func() []Output); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Output)
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
