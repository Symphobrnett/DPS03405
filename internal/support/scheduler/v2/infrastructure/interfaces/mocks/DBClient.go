// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	errors "github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddInterval provides a mock function with given fields: e
func (_m *DBClient) AddInterval(e models.Interval) (models.Interval, errors.EdgeX) {
	ret := _m.Called(e)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(models.Interval) models.Interval); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(models.Interval) errors.EdgeX); ok {
		r1 = rf(e)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}
