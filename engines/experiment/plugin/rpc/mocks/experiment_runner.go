// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	experimentrunner "github.com/caraml-dev/turing/engines/experiment/runner"

	json "encoding/json"

	metrics "github.com/gojek/mlp/api/pkg/instrumentation/metrics"

	mock "github.com/stretchr/testify/mock"
)

// ConfigurableExperimentRunner is an autogenerated mock type for the ConfigurableExperimentRunner type
type ConfigurableExperimentRunner struct {
	mock.Mock
}

// Configure provides a mock function with given fields: cfg
func (_m *ConfigurableExperimentRunner) Configure(cfg json.RawMessage) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(json.RawMessage) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTreatmentForRequest provides a mock function with given fields: header, payload, options
func (_m *ConfigurableExperimentRunner) GetTreatmentForRequest(header http.Header, payload []byte, options experimentrunner.GetTreatmentOptions) (*experimentrunner.Treatment, error) {
	ret := _m.Called(header, payload, options)

	var r0 *experimentrunner.Treatment
	if rf, ok := ret.Get(0).(func(http.Header, []byte, experimentrunner.GetTreatmentOptions) *experimentrunner.Treatment); ok {
		r0 = rf(header, payload, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*experimentrunner.Treatment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(http.Header, []byte, experimentrunner.GetTreatmentOptions) error); ok {
		r1 = rf(header, payload, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterMetricsCollector provides a mock function with given fields: collector, metricsRegistrationHelper
func (_m *ConfigurableExperimentRunner) RegisterMetricsCollector(collector metrics.Collector, metricsRegistrationHelper experimentrunner.MetricsRegistrationHelper) error {
	ret := _m.Called(collector, metricsRegistrationHelper)

	var r0 error
	if rf, ok := ret.Get(0).(func(metrics.Collector, experimentrunner.MetricsRegistrationHelper) error); ok {
		r0 = rf(collector, metricsRegistrationHelper)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewConfigurableExperimentRunner interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfigurableExperimentRunner creates a new instance of ConfigurableExperimentRunner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfigurableExperimentRunner(t mockConstructorTestingTNewConfigurableExperimentRunner) *ConfigurableExperimentRunner {
	mock := &ConfigurableExperimentRunner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
