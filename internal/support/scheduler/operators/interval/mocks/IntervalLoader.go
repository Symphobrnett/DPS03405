// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/models"

// IntervalLoader is an autogenerated mock type for the IntervalLoader type
type IntervalLoader struct {
	mock.Mock
}

// IntervalById provides a mock function with given fields: id
func (_m *IntervalLoader) IntervalById(id string) (models.Interval, error) {
	ret := _m.Called(id)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalByName provides a mock function with given fields: name
func (_m *IntervalLoader) IntervalByName(name string) (models.Interval, error) {
	ret := _m.Called(name)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
