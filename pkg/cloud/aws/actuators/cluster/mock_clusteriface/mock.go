// Copyright © 2018 The Kubernetes Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset/typed/cluster/v1alpha1 (interfaces: ClusterInterface)

// Package mock_clusteriface is a generated GoMock package.
package mock_clusteriface

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	reflect "reflect"
	v1alpha1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

// MockClusterInterface is a mock of ClusterInterface interface
type MockClusterInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterInterfaceMockRecorder
}

// MockClusterInterfaceMockRecorder is the mock recorder for MockClusterInterface
type MockClusterInterfaceMockRecorder struct {
	mock *MockClusterInterface
}

// NewMockClusterInterface creates a new mock instance
func NewMockClusterInterface(ctrl *gomock.Controller) *MockClusterInterface {
	mock := &MockClusterInterface{ctrl: ctrl}
	mock.recorder = &MockClusterInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterInterface) EXPECT() *MockClusterInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockClusterInterface) Create(arg0 *v1alpha1.Cluster) (*v1alpha1.Cluster, error) {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockClusterInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockClusterInterface) Delete(arg0 string, arg1 *v1.DeleteOptions) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockClusterInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *MockClusterInterface) DeleteCollection(arg0 *v1.DeleteOptions, arg1 v1.ListOptions) error {
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockClusterInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockClusterInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockClusterInterface) Get(arg0 string, arg1 v1.GetOptions) (*v1alpha1.Cluster, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClusterInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterInterface)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockClusterInterface) List(arg0 v1.ListOptions) (*v1alpha1.ClusterList, error) {
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1alpha1.ClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockClusterInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockClusterInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1alpha1.Cluster, error) {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockClusterInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockClusterInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockClusterInterface) Update(arg0 *v1alpha1.Cluster) (*v1alpha1.Cluster, error) {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockClusterInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterInterface)(nil).Update), arg0)
}

// UpdateStatus mocks base method
func (m *MockClusterInterface) UpdateStatus(arg0 *v1alpha1.Cluster) (*v1alpha1.Cluster, error) {
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockClusterInterfaceMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockClusterInterface)(nil).UpdateStatus), arg0)
}

// Watch mocks base method
func (m *MockClusterInterface) Watch(arg0 v1.ListOptions) (watch.Interface, error) {
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockClusterInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterInterface)(nil).Watch), arg0)
}
