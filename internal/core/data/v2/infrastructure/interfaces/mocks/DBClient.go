// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	errors "github.com/edgexfoundry/go-mod-core-contracts/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddEvent provides a mock function with given fields: e
func (_m *DBClient) AddEvent(e models.Event) (models.Event, errors.EdgeX) {
	ret := _m.Called(e)

	var r0 models.Event
	if rf, ok := ret.Get(0).(func(models.Event) models.Event); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(models.Event) errors.EdgeX); ok {
		r1 = rf(e)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllEvents provides a mock function with given fields: offset, limit
func (_m *DBClient) AllEvents(offset int, limit int) ([]models.Event, errors.EdgeX) {
	ret := _m.Called(offset, limit)

	var r0 []models.Event
	if rf, ok := ret.Get(0).(func(int, int) []models.Event); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int) errors.EdgeX); ok {
		r1 = rf(offset, limit)
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

// DeleteEventById provides a mock function with given fields: id
func (_m *DBClient) DeleteEventById(id string) errors.EdgeX {
	ret := _m.Called(id)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeletePushedEvents provides a mock function with given fields:
func (_m *DBClient) DeletePushedEvents() errors.EdgeX {
	ret := _m.Called()

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func() errors.EdgeX); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// EventById provides a mock function with given fields: id
func (_m *DBClient) EventById(id string) (models.Event, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 models.Event
	if rf, ok := ret.Get(0).(func(string) models.Event); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventCountByDevice provides a mock function with given fields: deviceName
func (_m *DBClient) EventCountByDevice(deviceName string) (uint32, errors.EdgeX) {
	ret := _m.Called(deviceName)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(deviceName)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(deviceName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventTotalCount provides a mock function with given fields:
func (_m *DBClient) EventTotalCount() (uint32, errors.EdgeX) {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func() errors.EdgeX); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventsByDeviceName provides a mock function with given fields: offset, limit, name
func (_m *DBClient) EventsByDeviceName(offset int, limit int, name string) ([]models.Event, errors.EdgeX) {
	ret := _m.Called(offset, limit, name)

	var r0 []models.Event
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Event); ok {
		r0 = rf(offset, limit, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateEventPushedById provides a mock function with given fields: id
func (_m *DBClient) UpdateEventPushedById(id string) errors.EdgeX {
	ret := _m.Called(id)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}
