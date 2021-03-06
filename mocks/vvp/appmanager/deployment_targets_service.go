// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"

	mock "github.com/stretchr/testify/mock"
)

// DeploymentTargetsService is an autogenerated mock type for the DeploymentTargetsService type
type DeploymentTargetsService struct {
	mock.Mock
}

// CreateDeploymentTarget provides a mock function with given fields: ctx, namespaceName, depTarget
func (_m *DeploymentTargetsService) CreateDeploymentTarget(ctx context.Context, namespaceName string, depTarget appmanagerapi.DeploymentTarget) (*appmanagerapi.DeploymentTarget, error) {
	ret := _m.Called(ctx, namespaceName, depTarget)

	var r0 *appmanagerapi.DeploymentTarget
	if rf, ok := ret.Get(0).(func(context.Context, string, appmanagerapi.DeploymentTarget) *appmanagerapi.DeploymentTarget); ok {
		r0 = rf(ctx, namespaceName, depTarget)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appmanagerapi.DeploymentTarget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, appmanagerapi.DeploymentTarget) error); ok {
		r1 = rf(ctx, namespaceName, depTarget)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDeploymentTarget provides a mock function with given fields: ctx, namespaceName, name
func (_m *DeploymentTargetsService) DeleteDeploymentTarget(ctx context.Context, namespaceName string, name string) (*appmanagerapi.DeploymentTarget, error) {
	ret := _m.Called(ctx, namespaceName, name)

	var r0 *appmanagerapi.DeploymentTarget
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *appmanagerapi.DeploymentTarget); ok {
		r0 = rf(ctx, namespaceName, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appmanagerapi.DeploymentTarget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespaceName, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeploymentTarget provides a mock function with given fields: ctx, namespaceName, name
func (_m *DeploymentTargetsService) GetDeploymentTarget(ctx context.Context, namespaceName string, name string) (*appmanagerapi.DeploymentTarget, error) {
	ret := _m.Called(ctx, namespaceName, name)

	var r0 *appmanagerapi.DeploymentTarget
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *appmanagerapi.DeploymentTarget); ok {
		r0 = rf(ctx, namespaceName, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appmanagerapi.DeploymentTarget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespaceName, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
