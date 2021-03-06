// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	models "github.com/gojek/merlin/models"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Delete provides a mock function with given fields: modelService
func (_m *Controller) Delete(modelService *models.Service) (*models.Service, error) {
	ret := _m.Called(modelService)

	var r0 *models.Service
	if rf, ok := ret.Get(0).(func(*models.Service) *models.Service); ok {
		r0 = rf(modelService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Service) error); ok {
		r1 = rf(modelService)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deploy provides a mock function with given fields: modelService
func (_m *Controller) Deploy(modelService *models.Service) (*models.Service, error) {
	ret := _m.Called(modelService)

	var r0 *models.Service
	if rf, ok := ret.Get(0).(func(*models.Service) *models.Service); ok {
		r0 = rf(modelService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Service) error); ok {
		r1 = rf(modelService)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContainers provides a mock function with given fields: namespace, labelSelector
func (_m *Controller) GetContainers(namespace string, labelSelector string) ([]*models.Container, error) {
	ret := _m.Called(namespace, labelSelector)

	var r0 []*models.Container
	if rf, ok := ret.Get(0).(func(string, string) []*models.Container); ok {
		r0 = rf(namespace, labelSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, labelSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
