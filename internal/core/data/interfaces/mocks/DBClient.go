// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import go_mod_core_contractsmodels "github.com/edgexfoundry/go-mod-core-contracts/models"

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/edgex-go/internal/pkg/correlation/models"

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddEvent provides a mock function with given fields: e
func (_m *DBClient) AddEvent(e models.Event) (string, error) {
	ret := _m.Called(e)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Event) string); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Event) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddReading provides a mock function with given fields: r
func (_m *DBClient) AddReading(r go_mod_core_contractsmodels.Reading) (string, error) {
	ret := _m.Called(r)

	var r0 string
	if rf, ok := ret.Get(0).(func(go_mod_core_contractsmodels.Reading) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(go_mod_core_contractsmodels.Reading) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddValueDescriptor provides a mock function with given fields: v
func (_m *DBClient) AddValueDescriptor(v go_mod_core_contractsmodels.ValueDescriptor) (string, error) {
	ret := _m.Called(v)

	var r0 string
	if rf, ok := ret.Get(0).(func(go_mod_core_contractsmodels.ValueDescriptor) string); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(go_mod_core_contractsmodels.ValueDescriptor) error); ok {
		r1 = rf(v)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// DeleteEventById provides a mock function with given fields: id
func (_m *DBClient) DeleteEventById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEventsByDevice provides a mock function with given fields: deviceId
func (_m *DBClient) DeleteEventsByDevice(deviceId string) (int, error) {
	ret := _m.Called(deviceId)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(deviceId)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deviceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteReadingById provides a mock function with given fields: id
func (_m *DBClient) DeleteReadingById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReadingsByDevice provides a mock function with given fields: device
func (_m *DBClient) DeleteReadingsByDevice(device string) error {
	ret := _m.Called(device)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(device)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteValueDescriptorById provides a mock function with given fields: id
func (_m *DBClient) DeleteValueDescriptorById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventById provides a mock function with given fields: id
func (_m *DBClient) EventById(id string) (go_mod_core_contractsmodels.Event, error) {
	ret := _m.Called(id)

	var r0 go_mod_core_contractsmodels.Event
	if rf, ok := ret.Get(0).(func(string) go_mod_core_contractsmodels.Event); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(go_mod_core_contractsmodels.Event)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventCount provides a mock function with given fields:
func (_m *DBClient) EventCount() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventCountByDeviceId provides a mock function with given fields: id
func (_m *DBClient) EventCountByDeviceId(id string) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Events provides a mock function with given fields:
func (_m *DBClient) Events() ([]go_mod_core_contractsmodels.Event, error) {
	ret := _m.Called()

	var r0 []go_mod_core_contractsmodels.Event
	if rf, ok := ret.Get(0).(func() []go_mod_core_contractsmodels.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Event)
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

// EventsByChecksum provides a mock function with given fields: checksum
func (_m *DBClient) EventsByChecksum(checksum string) ([]go_mod_core_contractsmodels.Event, error) {
	ret := _m.Called(checksum)

	var r0 []go_mod_core_contractsmodels.Event
	if rf, ok := ret.Get(0).(func(string) []go_mod_core_contractsmodels.Event); ok {
		r0 = rf(checksum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(checksum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsByCreationTime provides a mock function with given fields: startTime, endTime, limit
func (_m *DBClient) EventsByCreationTime(startTime int64, endTime int64, limit int) ([]go_mod_core_contractsmodels.Event, error) {
	ret := _m.Called(startTime, endTime, limit)

	var r0 []go_mod_core_contractsmodels.Event
	if rf, ok := ret.Get(0).(func(int64, int64, int) []go_mod_core_contractsmodels.Event); ok {
		r0 = rf(startTime, endTime, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, int) error); ok {
		r1 = rf(startTime, endTime, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsForDevice provides a mock function with given fields: id
func (_m *DBClient) EventsForDevice(id string) ([]go_mod_core_contractsmodels.Event, error) {
	ret := _m.Called(id)

	var r0 []go_mod_core_contractsmodels.Event
	if rf, ok := ret.Get(0).(func(string) []go_mod_core_contractsmodels.Event); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsForDeviceLimit provides a mock function with given fields: id, limit
func (_m *DBClient) EventsForDeviceLimit(id string, limit int) ([]go_mod_core_contractsmodels.Event, error) {
	ret := _m.Called(id, limit)

	var r0 []go_mod_core_contractsmodels.Event
	if rf, ok := ret.Get(0).(func(string, int) []go_mod_core_contractsmodels.Event); ok {
		r0 = rf(id, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(id, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsOlderThanAge provides a mock function with given fields: age
func (_m *DBClient) EventsOlderThanAge(age int64) ([]go_mod_core_contractsmodels.Event, error) {
	ret := _m.Called(age)

	var r0 []go_mod_core_contractsmodels.Event
	if rf, ok := ret.Get(0).(func(int64) []go_mod_core_contractsmodels.Event); ok {
		r0 = rf(age)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(age)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsPushed provides a mock function with given fields:
func (_m *DBClient) EventsPushed() ([]go_mod_core_contractsmodels.Event, error) {
	ret := _m.Called()

	var r0 []go_mod_core_contractsmodels.Event
	if rf, ok := ret.Get(0).(func() []go_mod_core_contractsmodels.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Event)
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

// EventsWithLimit provides a mock function with given fields: limit
func (_m *DBClient) EventsWithLimit(limit int) ([]go_mod_core_contractsmodels.Event, error) {
	ret := _m.Called(limit)

	var r0 []go_mod_core_contractsmodels.Event
	if rf, ok := ret.Get(0).(func(int) []go_mod_core_contractsmodels.Event); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingById provides a mock function with given fields: id
func (_m *DBClient) ReadingById(id string) (go_mod_core_contractsmodels.Reading, error) {
	ret := _m.Called(id)

	var r0 go_mod_core_contractsmodels.Reading
	if rf, ok := ret.Get(0).(func(string) go_mod_core_contractsmodels.Reading); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(go_mod_core_contractsmodels.Reading)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingCount provides a mock function with given fields:
func (_m *DBClient) ReadingCount() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readings provides a mock function with given fields:
func (_m *DBClient) Readings() ([]go_mod_core_contractsmodels.Reading, error) {
	ret := _m.Called()

	var r0 []go_mod_core_contractsmodels.Reading
	if rf, ok := ret.Get(0).(func() []go_mod_core_contractsmodels.Reading); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Reading)
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

// ReadingsByCreationTime provides a mock function with given fields: start, end, limit
func (_m *DBClient) ReadingsByCreationTime(start int64, end int64, limit int) ([]go_mod_core_contractsmodels.Reading, error) {
	ret := _m.Called(start, end, limit)

	var r0 []go_mod_core_contractsmodels.Reading
	if rf, ok := ret.Get(0).(func(int64, int64, int) []go_mod_core_contractsmodels.Reading); ok {
		r0 = rf(start, end, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, int) error); ok {
		r1 = rf(start, end, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsByDevice provides a mock function with given fields: id, limit
func (_m *DBClient) ReadingsByDevice(id string, limit int) ([]go_mod_core_contractsmodels.Reading, error) {
	ret := _m.Called(id, limit)

	var r0 []go_mod_core_contractsmodels.Reading
	if rf, ok := ret.Get(0).(func(string, int) []go_mod_core_contractsmodels.Reading); ok {
		r0 = rf(id, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(id, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsByDeviceAndValueDescriptor provides a mock function with given fields: deviceId, valueDescriptor, limit
func (_m *DBClient) ReadingsByDeviceAndValueDescriptor(deviceId string, valueDescriptor string, limit int) ([]go_mod_core_contractsmodels.Reading, error) {
	ret := _m.Called(deviceId, valueDescriptor, limit)

	var r0 []go_mod_core_contractsmodels.Reading
	if rf, ok := ret.Get(0).(func(string, string, int) []go_mod_core_contractsmodels.Reading); ok {
		r0 = rf(deviceId, valueDescriptor, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int) error); ok {
		r1 = rf(deviceId, valueDescriptor, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsByValueDescriptor provides a mock function with given fields: name, limit
func (_m *DBClient) ReadingsByValueDescriptor(name string, limit int) ([]go_mod_core_contractsmodels.Reading, error) {
	ret := _m.Called(name, limit)

	var r0 []go_mod_core_contractsmodels.Reading
	if rf, ok := ret.Get(0).(func(string, int) []go_mod_core_contractsmodels.Reading); ok {
		r0 = rf(name, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(name, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsByValueDescriptorNames provides a mock function with given fields: names, limit
func (_m *DBClient) ReadingsByValueDescriptorNames(names []string, limit int) ([]go_mod_core_contractsmodels.Reading, error) {
	ret := _m.Called(names, limit)

	var r0 []go_mod_core_contractsmodels.Reading
	if rf, ok := ret.Get(0).(func([]string, int) []go_mod_core_contractsmodels.Reading); ok {
		r0 = rf(names, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, int) error); ok {
		r1 = rf(names, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScrubAllEvents provides a mock function with given fields:
func (_m *DBClient) ScrubAllEvents() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ScrubAllValueDescriptors provides a mock function with given fields:
func (_m *DBClient) ScrubAllValueDescriptors() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEvent provides a mock function with given fields: e
func (_m *DBClient) UpdateEvent(e models.Event) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Event) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateReading provides a mock function with given fields: r
func (_m *DBClient) UpdateReading(r go_mod_core_contractsmodels.Reading) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(go_mod_core_contractsmodels.Reading) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateValueDescriptor provides a mock function with given fields: v
func (_m *DBClient) UpdateValueDescriptor(v go_mod_core_contractsmodels.ValueDescriptor) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(go_mod_core_contractsmodels.ValueDescriptor) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValueDescriptorById provides a mock function with given fields: id
func (_m *DBClient) ValueDescriptorById(id string) (go_mod_core_contractsmodels.ValueDescriptor, error) {
	ret := _m.Called(id)

	var r0 go_mod_core_contractsmodels.ValueDescriptor
	if rf, ok := ret.Get(0).(func(string) go_mod_core_contractsmodels.ValueDescriptor); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(go_mod_core_contractsmodels.ValueDescriptor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValueDescriptorByName provides a mock function with given fields: name
func (_m *DBClient) ValueDescriptorByName(name string) (go_mod_core_contractsmodels.ValueDescriptor, error) {
	ret := _m.Called(name)

	var r0 go_mod_core_contractsmodels.ValueDescriptor
	if rf, ok := ret.Get(0).(func(string) go_mod_core_contractsmodels.ValueDescriptor); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(go_mod_core_contractsmodels.ValueDescriptor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValueDescriptors provides a mock function with given fields:
func (_m *DBClient) ValueDescriptors() ([]go_mod_core_contractsmodels.ValueDescriptor, error) {
	ret := _m.Called()

	var r0 []go_mod_core_contractsmodels.ValueDescriptor
	if rf, ok := ret.Get(0).(func() []go_mod_core_contractsmodels.ValueDescriptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.ValueDescriptor)
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

// ValueDescriptorsByLabel provides a mock function with given fields: label
func (_m *DBClient) ValueDescriptorsByLabel(label string) ([]go_mod_core_contractsmodels.ValueDescriptor, error) {
	ret := _m.Called(label)

	var r0 []go_mod_core_contractsmodels.ValueDescriptor
	if rf, ok := ret.Get(0).(func(string) []go_mod_core_contractsmodels.ValueDescriptor); ok {
		r0 = rf(label)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.ValueDescriptor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(label)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValueDescriptorsByName provides a mock function with given fields: names
func (_m *DBClient) ValueDescriptorsByName(names []string) ([]go_mod_core_contractsmodels.ValueDescriptor, error) {
	ret := _m.Called(names)

	var r0 []go_mod_core_contractsmodels.ValueDescriptor
	if rf, ok := ret.Get(0).(func([]string) []go_mod_core_contractsmodels.ValueDescriptor); ok {
		r0 = rf(names)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.ValueDescriptor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(names)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValueDescriptorsByType provides a mock function with given fields: t
func (_m *DBClient) ValueDescriptorsByType(t string) ([]go_mod_core_contractsmodels.ValueDescriptor, error) {
	ret := _m.Called(t)

	var r0 []go_mod_core_contractsmodels.ValueDescriptor
	if rf, ok := ret.Get(0).(func(string) []go_mod_core_contractsmodels.ValueDescriptor); ok {
		r0 = rf(t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.ValueDescriptor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValueDescriptorsByUomLabel provides a mock function with given fields: uomLabel
func (_m *DBClient) ValueDescriptorsByUomLabel(uomLabel string) ([]go_mod_core_contractsmodels.ValueDescriptor, error) {
	ret := _m.Called(uomLabel)

	var r0 []go_mod_core_contractsmodels.ValueDescriptor
	if rf, ok := ret.Get(0).(func(string) []go_mod_core_contractsmodels.ValueDescriptor); ok {
		r0 = rf(uomLabel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_mod_core_contractsmodels.ValueDescriptor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uomLabel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
