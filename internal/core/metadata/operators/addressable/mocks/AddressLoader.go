// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/models"

// AddressLoader is an autogenerated mock type for the AddressLoader type
type AddressLoader struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *AddressLoader) Execute() ([]models.Addressable, error) {
	ret := _m.Called()

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func() []models.Addressable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
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
