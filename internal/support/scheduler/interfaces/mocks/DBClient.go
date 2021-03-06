// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/edgex-go/pkg/models"

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddInterval provides a mock function with given fields: interval
func (_m *DBClient) AddInterval(interval models.Interval) (string, error) {
	ret := _m.Called(interval)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Interval) string); ok {
		r0 = rf(interval)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Interval) error); ok {
		r1 = rf(interval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddIntervalAction provides a mock function with given fields: intervalAction
func (_m *DBClient) AddIntervalAction(intervalAction models.IntervalAction) (string, error) {
	ret := _m.Called(intervalAction)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.IntervalAction) string); ok {
		r0 = rf(intervalAction)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.IntervalAction) error); ok {
		r1 = rf(intervalAction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// Connect provides a mock function with given fields:
func (_m *DBClient) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteIntervalActionById provides a mock function with given fields: id
func (_m *DBClient) DeleteIntervalActionById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteIntervalById provides a mock function with given fields: id
func (_m *DBClient) DeleteIntervalById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IntervalActionById provides a mock function with given fields: id
func (_m *DBClient) IntervalActionById(id string) (models.IntervalAction, error) {
	ret := _m.Called(id)

	var r0 models.IntervalAction
	if rf, ok := ret.Get(0).(func(string) models.IntervalAction); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.IntervalAction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalActionByName provides a mock function with given fields: name
func (_m *DBClient) IntervalActionByName(name string) (models.IntervalAction, error) {
	ret := _m.Called(name)

	var r0 models.IntervalAction
	if rf, ok := ret.Get(0).(func(string) models.IntervalAction); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.IntervalAction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalActions provides a mock function with given fields:
func (_m *DBClient) IntervalActions() ([]models.IntervalAction, error) {
	ret := _m.Called()

	var r0 []models.IntervalAction
	if rf, ok := ret.Get(0).(func() []models.IntervalAction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.IntervalAction)
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

// IntervalActionsByIntervalName provides a mock function with given fields: name
func (_m *DBClient) IntervalActionsByIntervalName(name string) ([]models.IntervalAction, error) {
	ret := _m.Called(name)

	var r0 []models.IntervalAction
	if rf, ok := ret.Get(0).(func(string) []models.IntervalAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.IntervalAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalActionsByTarget provides a mock function with given fields: name
func (_m *DBClient) IntervalActionsByTarget(name string) ([]models.IntervalAction, error) {
	ret := _m.Called(name)

	var r0 []models.IntervalAction
	if rf, ok := ret.Get(0).(func(string) []models.IntervalAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.IntervalAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalActionsWithLimit provides a mock function with given fields: limit
func (_m *DBClient) IntervalActionsWithLimit(limit int) ([]models.IntervalAction, error) {
	ret := _m.Called(limit)

	var r0 []models.IntervalAction
	if rf, ok := ret.Get(0).(func(int) []models.IntervalAction); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.IntervalAction)
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

// IntervalById provides a mock function with given fields: id
func (_m *DBClient) IntervalById(id string) (models.Interval, error) {
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
func (_m *DBClient) IntervalByName(name string) (models.Interval, error) {
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

// Intervals provides a mock function with given fields:
func (_m *DBClient) Intervals() ([]models.Interval, error) {
	ret := _m.Called()

	var r0 []models.Interval
	if rf, ok := ret.Get(0).(func() []models.Interval); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Interval)
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

// IntervalsWithLimit provides a mock function with given fields: limit
func (_m *DBClient) IntervalsWithLimit(limit int) ([]models.Interval, error) {
	ret := _m.Called(limit)

	var r0 []models.Interval
	if rf, ok := ret.Get(0).(func(int) []models.Interval); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Interval)
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

// ScrubAllIntervalActions provides a mock function with given fields:
func (_m *DBClient) ScrubAllIntervalActions() (int, error) {
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

// ScrubAllIntervals provides a mock function with given fields:
func (_m *DBClient) ScrubAllIntervals() (int, error) {
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

// UpdateInterval provides a mock function with given fields: interval
func (_m *DBClient) UpdateInterval(interval models.Interval) error {
	ret := _m.Called(interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Interval) error); ok {
		r0 = rf(interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateIntervalAction provides a mock function with given fields: intervalAction
func (_m *DBClient) UpdateIntervalAction(intervalAction models.IntervalAction) error {
	ret := _m.Called(intervalAction)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.IntervalAction) error); ok {
		r0 = rf(intervalAction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
