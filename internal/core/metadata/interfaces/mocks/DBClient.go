// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/models"

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddAddressable provides a mock function with given fields: a
func (_m *DBClient) AddAddressable(a models.Addressable) (string, error) {
	ret := _m.Called(a)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Addressable) string); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Addressable) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddCommand provides a mock function with given fields: c
func (_m *DBClient) AddCommand(c models.Command) (string, error) {
	ret := _m.Called(c)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Command) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Command) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddDevice provides a mock function with given fields: d
func (_m *DBClient) AddDevice(d models.Device) (string, error) {
	ret := _m.Called(d)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Device) string); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Device) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddDeviceProfile provides a mock function with given fields: d
func (_m *DBClient) AddDeviceProfile(d models.DeviceProfile) (string, error) {
	ret := _m.Called(d)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.DeviceProfile) string); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.DeviceProfile) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddDeviceReport provides a mock function with given fields: dr
func (_m *DBClient) AddDeviceReport(dr models.DeviceReport) (string, error) {
	ret := _m.Called(dr)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.DeviceReport) string); ok {
		r0 = rf(dr)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.DeviceReport) error); ok {
		r1 = rf(dr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddDeviceService provides a mock function with given fields: ds
func (_m *DBClient) AddDeviceService(ds models.DeviceService) (string, error) {
	ret := _m.Called(ds)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.DeviceService) string); ok {
		r0 = rf(ds)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.DeviceService) error); ok {
		r1 = rf(ds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddProvisionWatcher provides a mock function with given fields: pw
func (_m *DBClient) AddProvisionWatcher(pw models.ProvisionWatcher) (string, error) {
	ret := _m.Called(pw)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.ProvisionWatcher) string); ok {
		r0 = rf(pw)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.ProvisionWatcher) error); ok {
		r1 = rf(pw)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// DeleteAddressableById provides a mock function with given fields: id
func (_m *DBClient) DeleteAddressableById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCommandById provides a mock function with given fields: id
func (_m *DBClient) DeleteCommandById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeviceById provides a mock function with given fields: id
func (_m *DBClient) DeleteDeviceById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeviceProfileById provides a mock function with given fields: id
func (_m *DBClient) DeleteDeviceProfileById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeviceReportById provides a mock function with given fields: id
func (_m *DBClient) DeleteDeviceReportById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeviceServiceById provides a mock function with given fields: id
func (_m *DBClient) DeleteDeviceServiceById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProvisionWatcherById provides a mock function with given fields: id
func (_m *DBClient) DeleteProvisionWatcherById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAddressableById provides a mock function with given fields: id
func (_m *DBClient) GetAddressableById(id string) (models.Addressable, error) {
	ret := _m.Called(id)

	var r0 models.Addressable
	if rf, ok := ret.Get(0).(func(string) models.Addressable); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Addressable)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressableByName provides a mock function with given fields: n
func (_m *DBClient) GetAddressableByName(n string) (models.Addressable, error) {
	ret := _m.Called(n)

	var r0 models.Addressable
	if rf, ok := ret.Get(0).(func(string) models.Addressable); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.Addressable)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressables provides a mock function with given fields:
func (_m *DBClient) GetAddressables() ([]models.Addressable, error) {
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

// GetAddressablesByAddress provides a mock function with given fields: add
func (_m *DBClient) GetAddressablesByAddress(add string) ([]models.Addressable, error) {
	ret := _m.Called(add)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(string) []models.Addressable); ok {
		r0 = rf(add)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(add)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByPort provides a mock function with given fields: p
func (_m *DBClient) GetAddressablesByPort(p int) ([]models.Addressable, error) {
	ret := _m.Called(p)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(int) []models.Addressable); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByPublisher provides a mock function with given fields: p
func (_m *DBClient) GetAddressablesByPublisher(p string) ([]models.Addressable, error) {
	ret := _m.Called(p)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(string) []models.Addressable); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByTopic provides a mock function with given fields: t
func (_m *DBClient) GetAddressablesByTopic(t string) ([]models.Addressable, error) {
	ret := _m.Called(t)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(string) []models.Addressable); ok {
		r0 = rf(t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
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

// GetAllCommands provides a mock function with given fields:
func (_m *DBClient) GetAllCommands() ([]models.Command, error) {
	ret := _m.Called()

	var r0 []models.Command
	if rf, ok := ret.Get(0).(func() []models.Command); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Command)
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

// GetAllDeviceProfiles provides a mock function with given fields:
func (_m *DBClient) GetAllDeviceProfiles() ([]models.DeviceProfile, error) {
	ret := _m.Called()

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func() []models.DeviceProfile); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
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

// GetAllDeviceReports provides a mock function with given fields:
func (_m *DBClient) GetAllDeviceReports() ([]models.DeviceReport, error) {
	ret := _m.Called()

	var r0 []models.DeviceReport
	if rf, ok := ret.Get(0).(func() []models.DeviceReport); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceReport)
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

// GetAllDeviceServices provides a mock function with given fields:
func (_m *DBClient) GetAllDeviceServices() ([]models.DeviceService, error) {
	ret := _m.Called()

	var r0 []models.DeviceService
	if rf, ok := ret.Get(0).(func() []models.DeviceService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceService)
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

// GetAllDevices provides a mock function with given fields:
func (_m *DBClient) GetAllDevices() ([]models.Device, error) {
	ret := _m.Called()

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func() []models.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
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

// GetAllProvisionWatchers provides a mock function with given fields:
func (_m *DBClient) GetAllProvisionWatchers() ([]models.ProvisionWatcher, error) {
	ret := _m.Called()

	var r0 []models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func() []models.ProvisionWatcher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProvisionWatcher)
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

// GetCommandById provides a mock function with given fields: id
func (_m *DBClient) GetCommandById(id string) (models.Command, error) {
	ret := _m.Called(id)

	var r0 models.Command
	if rf, ok := ret.Get(0).(func(string) models.Command); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Command)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommandByName provides a mock function with given fields: id
func (_m *DBClient) GetCommandByName(id string) ([]models.Command, error) {
	ret := _m.Called(id)

	var r0 []models.Command
	if rf, ok := ret.Get(0).(func(string) []models.Command); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Command)
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

// GetDeviceById provides a mock function with given fields: id
func (_m *DBClient) GetDeviceById(id string) (models.Device, error) {
	ret := _m.Called(id)

	var r0 models.Device
	if rf, ok := ret.Get(0).(func(string) models.Device); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceByName provides a mock function with given fields: n
func (_m *DBClient) GetDeviceByName(n string) (models.Device, error) {
	ret := _m.Called(n)

	var r0 models.Device
	if rf, ok := ret.Get(0).(func(string) models.Device); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfileById provides a mock function with given fields: id
func (_m *DBClient) GetDeviceProfileById(id string) (models.DeviceProfile, error) {
	ret := _m.Called(id)

	var r0 models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) models.DeviceProfile); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.DeviceProfile)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfileByName provides a mock function with given fields: n
func (_m *DBClient) GetDeviceProfileByName(n string) (models.DeviceProfile, error) {
	ret := _m.Called(n)

	var r0 models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) models.DeviceProfile); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.DeviceProfile)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfilesByManufacturer provides a mock function with given fields: man
func (_m *DBClient) GetDeviceProfilesByManufacturer(man string) ([]models.DeviceProfile, error) {
	ret := _m.Called(man)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) []models.DeviceProfile); ok {
		r0 = rf(man)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(man)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfilesByManufacturerModel provides a mock function with given fields: man, mod
func (_m *DBClient) GetDeviceProfilesByManufacturerModel(man string, mod string) ([]models.DeviceProfile, error) {
	ret := _m.Called(man, mod)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string, string) []models.DeviceProfile); ok {
		r0 = rf(man, mod)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(man, mod)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfilesByModel provides a mock function with given fields: m
func (_m *DBClient) GetDeviceProfilesByModel(m string) ([]models.DeviceProfile, error) {
	ret := _m.Called(m)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) []models.DeviceProfile); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfilesWithLabel provides a mock function with given fields: l
func (_m *DBClient) GetDeviceProfilesWithLabel(l string) ([]models.DeviceProfile, error) {
	ret := _m.Called(l)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) []models.DeviceProfile); ok {
		r0 = rf(l)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(l)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceReportByDeviceName provides a mock function with given fields: n
func (_m *DBClient) GetDeviceReportByDeviceName(n string) ([]models.DeviceReport, error) {
	ret := _m.Called(n)

	var r0 []models.DeviceReport
	if rf, ok := ret.Get(0).(func(string) []models.DeviceReport); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceReportById provides a mock function with given fields: id
func (_m *DBClient) GetDeviceReportById(id string) (models.DeviceReport, error) {
	ret := _m.Called(id)

	var r0 models.DeviceReport
	if rf, ok := ret.Get(0).(func(string) models.DeviceReport); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.DeviceReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceReportByName provides a mock function with given fields: n
func (_m *DBClient) GetDeviceReportByName(n string) (models.DeviceReport, error) {
	ret := _m.Called(n)

	var r0 models.DeviceReport
	if rf, ok := ret.Get(0).(func(string) models.DeviceReport); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.DeviceReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceReportsByAction provides a mock function with given fields: n
func (_m *DBClient) GetDeviceReportsByAction(n string) ([]models.DeviceReport, error) {
	ret := _m.Called(n)

	var r0 []models.DeviceReport
	if rf, ok := ret.Get(0).(func(string) []models.DeviceReport); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceServiceById provides a mock function with given fields: id
func (_m *DBClient) GetDeviceServiceById(id string) (models.DeviceService, error) {
	ret := _m.Called(id)

	var r0 models.DeviceService
	if rf, ok := ret.Get(0).(func(string) models.DeviceService); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.DeviceService)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceServiceByName provides a mock function with given fields: n
func (_m *DBClient) GetDeviceServiceByName(n string) (models.DeviceService, error) {
	ret := _m.Called(n)

	var r0 models.DeviceService
	if rf, ok := ret.Get(0).(func(string) models.DeviceService); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.DeviceService)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceServicesByAddressableId provides a mock function with given fields: id
func (_m *DBClient) GetDeviceServicesByAddressableId(id string) ([]models.DeviceService, error) {
	ret := _m.Called(id)

	var r0 []models.DeviceService
	if rf, ok := ret.Get(0).(func(string) []models.DeviceService); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceService)
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

// GetDeviceServicesWithLabel provides a mock function with given fields: l
func (_m *DBClient) GetDeviceServicesWithLabel(l string) ([]models.DeviceService, error) {
	ret := _m.Called(l)

	var r0 []models.DeviceService
	if rf, ok := ret.Get(0).(func(string) []models.DeviceService); ok {
		r0 = rf(l)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceService)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(l)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevicesByProfileId provides a mock function with given fields: pid
func (_m *DBClient) GetDevicesByProfileId(pid string) ([]models.Device, error) {
	ret := _m.Called(pid)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string) []models.Device); ok {
		r0 = rf(pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevicesByServiceId provides a mock function with given fields: sid
func (_m *DBClient) GetDevicesByServiceId(sid string) ([]models.Device, error) {
	ret := _m.Called(sid)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string) []models.Device); ok {
		r0 = rf(sid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevicesWithLabel provides a mock function with given fields: l
func (_m *DBClient) GetDevicesWithLabel(l string) ([]models.Device, error) {
	ret := _m.Called(l)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string) []models.Device); ok {
		r0 = rf(l)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(l)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionWatcherById provides a mock function with given fields: id
func (_m *DBClient) GetProvisionWatcherById(id string) (models.ProvisionWatcher, error) {
	ret := _m.Called(id)

	var r0 models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(string) models.ProvisionWatcher); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.ProvisionWatcher)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionWatcherByName provides a mock function with given fields: n
func (_m *DBClient) GetProvisionWatcherByName(n string) (models.ProvisionWatcher, error) {
	ret := _m.Called(n)

	var r0 models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(string) models.ProvisionWatcher); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.ProvisionWatcher)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionWatchersByIdentifier provides a mock function with given fields: k, v
func (_m *DBClient) GetProvisionWatchersByIdentifier(k string, v string) ([]models.ProvisionWatcher, error) {
	ret := _m.Called(k, v)

	var r0 []models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(string, string) []models.ProvisionWatcher); ok {
		r0 = rf(k, v)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProvisionWatcher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(k, v)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionWatchersByProfileId provides a mock function with given fields: id
func (_m *DBClient) GetProvisionWatchersByProfileId(id string) ([]models.ProvisionWatcher, error) {
	ret := _m.Called(id)

	var r0 []models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(string) []models.ProvisionWatcher); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProvisionWatcher)
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

// GetProvisionWatchersByServiceId provides a mock function with given fields: id
func (_m *DBClient) GetProvisionWatchersByServiceId(id string) ([]models.ProvisionWatcher, error) {
	ret := _m.Called(id)

	var r0 []models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(string) []models.ProvisionWatcher); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProvisionWatcher)
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

// ScrubMetadata provides a mock function with given fields:
func (_m *DBClient) ScrubMetadata() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAddressable provides a mock function with given fields: a
func (_m *DBClient) UpdateAddressable(a models.Addressable) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Addressable) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCommand provides a mock function with given fields: c
func (_m *DBClient) UpdateCommand(c models.Command) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Command) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDevice provides a mock function with given fields: d
func (_m *DBClient) UpdateDevice(d models.Device) error {
	ret := _m.Called(d)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Device) error); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceProfile provides a mock function with given fields: dp
func (_m *DBClient) UpdateDeviceProfile(dp models.DeviceProfile) error {
	ret := _m.Called(dp)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.DeviceProfile) error); ok {
		r0 = rf(dp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceReport provides a mock function with given fields: dr
func (_m *DBClient) UpdateDeviceReport(dr models.DeviceReport) error {
	ret := _m.Called(dr)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.DeviceReport) error); ok {
		r0 = rf(dr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceService provides a mock function with given fields: ds
func (_m *DBClient) UpdateDeviceService(ds models.DeviceService) error {
	ret := _m.Called(ds)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.DeviceService) error); ok {
		r0 = rf(ds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProvisionWatcher provides a mock function with given fields: pw
func (_m *DBClient) UpdateProvisionWatcher(pw models.ProvisionWatcher) error {
	ret := _m.Called(pw)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.ProvisionWatcher) error); ok {
		r0 = rf(pw)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
