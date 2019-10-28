// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DataAccessLayer is an autogenerated mock type for the DataAccessLayer type
type DataAccessLayer struct {
	mock.Mock
}

// Insert provides a mock function with given fields: collectionName, docs
func (_m *DataAccessLayer) Insert(collectionName string, docs ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, collectionName)
	_ca = append(_ca, docs...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(collectionName, docs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
