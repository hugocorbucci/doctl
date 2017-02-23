package mocks

import "github.com/digitalocean/doctl/do"
import "github.com/digitalocean/godo"
import "github.com/stretchr/testify/mock"

// Generated: please do not edit by hand

// LoadBalancersService is an autogenerated mock type for the LoadBalancersService type
type LoadBalancersService struct {
	mock.Mock
}

// AddDroplets provides a mock function with given fields: lbID, dIDs
func (_m *LoadBalancersService) AddDroplets(lbID string, dIDs ...int) error {
	ret := _m.Called(lbID, dIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...int) error); ok {
		r0 = rf(lbID, dIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddForwardingRules provides a mock function with given fields: lbID, rules
func (_m *LoadBalancersService) AddForwardingRules(lbID string, rules ...godo.ForwardingRule) error {
	ret := _m.Called(lbID, rules)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...godo.ForwardingRule) error); ok {
		r0 = rf(lbID, rules...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: lbr
func (_m *LoadBalancersService) Create(lbr *godo.LoadBalancerRequest) (*do.LoadBalancer, error) {
	ret := _m.Called(lbr)

	var r0 *do.LoadBalancer
	if rf, ok := ret.Get(0).(func(*godo.LoadBalancerRequest) *do.LoadBalancer); ok {
		r0 = rf(lbr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.LoadBalancer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.LoadBalancerRequest) error); ok {
		r1 = rf(lbr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: lbID
func (_m *LoadBalancersService) Delete(lbID string) error {
	ret := _m.Called(lbID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(lbID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: lbID
func (_m *LoadBalancersService) Get(lbID string) (*do.LoadBalancer, error) {
	ret := _m.Called(lbID)

	var r0 *do.LoadBalancer
	if rf, ok := ret.Get(0).(func(string) *do.LoadBalancer); ok {
		r0 = rf(lbID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.LoadBalancer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(lbID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *LoadBalancersService) List() (do.LoadBalancers, error) {
	ret := _m.Called()

	var r0 do.LoadBalancers
	if rf, ok := ret.Get(0).(func() do.LoadBalancers); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.LoadBalancers)
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

// RemoveDroplets provides a mock function with given fields: lbID, dIDs
func (_m *LoadBalancersService) RemoveDroplets(lbID string, dIDs ...int) error {
	ret := _m.Called(lbID, dIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...int) error); ok {
		r0 = rf(lbID, dIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveForwardingRules provides a mock function with given fields: lbID, rules
func (_m *LoadBalancersService) RemoveForwardingRules(lbID string, rules ...godo.ForwardingRule) error {
	ret := _m.Called(lbID, rules)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...godo.ForwardingRule) error); ok {
		r0 = rf(lbID, rules...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: lbID, lbr
func (_m *LoadBalancersService) Update(lbID string, lbr *godo.LoadBalancerRequest) (*do.LoadBalancer, error) {
	ret := _m.Called(lbID, lbr)

	var r0 *do.LoadBalancer
	if rf, ok := ret.Get(0).(func(string, *godo.LoadBalancerRequest) *do.LoadBalancer); ok {
		r0 = rf(lbID, lbr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.LoadBalancer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.LoadBalancerRequest) error); ok {
		r1 = rf(lbID, lbr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ do.LoadBalancersService = (*LoadBalancersService)(nil)
