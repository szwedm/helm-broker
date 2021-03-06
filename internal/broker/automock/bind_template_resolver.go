// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import bind "github.com/kyma-project/helm-broker/internal/bind"

import internal "github.com/kyma-project/helm-broker/internal"
import mock "github.com/stretchr/testify/mock"

// bindTemplateResolver is an autogenerated mock type for the bindTemplateResolver type
type bindTemplateResolver struct {
	mock.Mock
}

// Resolve provides a mock function with given fields: bindYAML, ns
func (_m *bindTemplateResolver) Resolve(bindYAML bind.RenderedBindYAML, ns internal.Namespace) (*bind.ResolveOutput, error) {
	ret := _m.Called(bindYAML, ns)

	var r0 *bind.ResolveOutput
	if rf, ok := ret.Get(0).(func(bind.RenderedBindYAML, internal.Namespace) *bind.ResolveOutput); ok {
		r0 = rf(bindYAML, ns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bind.ResolveOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bind.RenderedBindYAML, internal.Namespace) error); ok {
		r1 = rf(bindYAML, ns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
