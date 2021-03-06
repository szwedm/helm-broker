// Code generated by mockery v1.0.0
package automock

import internal "github.com/kyma-project/helm-broker/internal"
import mock "github.com/stretchr/testify/mock"

// CommonClient is an autogenerated mock type for the CommonClient type
type CommonClient struct {
	mock.Mock
}

// IsNamespaceScoped provides a mock function with given fields:
func (_m *CommonClient) IsNamespaceScoped() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListConfigurations provides a mock function with given fields:
func (_m *CommonClient) ListConfigurations() ([]internal.CommonAddon, error) {
	ret := _m.Called()

	var r0 []internal.CommonAddon
	if rf, ok := ret.Get(0).(func() []internal.CommonAddon); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]internal.CommonAddon)
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

// ReprocessRequest provides a mock function with given fields: addonName
func (_m *CommonClient) ReprocessRequest(addonName string) error {
	ret := _m.Called(addonName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(addonName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetNamespace provides a mock function with given fields: namespace
func (_m *CommonClient) SetNamespace(namespace string) {
	_m.Called(namespace)
}

// UpdateConfiguration provides a mock function with given fields: _a0
func (_m *CommonClient) UpdateConfiguration(_a0 *internal.CommonAddon) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*internal.CommonAddon) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateConfigurationStatus provides a mock function with given fields: _a0
func (_m *CommonClient) UpdateConfigurationStatus(_a0 *internal.CommonAddon) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*internal.CommonAddon) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
