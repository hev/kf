// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/kf/pkg/kf/buildpacks/fake (interfaces: BuilderCreator)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// FakeBuilderCreator is a mock of BuilderCreator interface
type FakeBuilderCreator struct {
	ctrl     *gomock.Controller
	recorder *FakeBuilderCreatorMockRecorder
}

// FakeBuilderCreatorMockRecorder is the mock recorder for FakeBuilderCreator
type FakeBuilderCreatorMockRecorder struct {
	mock *FakeBuilderCreator
}

// NewFakeBuilderCreator creates a new mock instance
func NewFakeBuilderCreator(ctrl *gomock.Controller) *FakeBuilderCreator {
	mock := &FakeBuilderCreator{ctrl: ctrl}
	mock.recorder = &FakeBuilderCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeBuilderCreator) EXPECT() *FakeBuilderCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *FakeBuilderCreator) Create(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *FakeBuilderCreatorMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*FakeBuilderCreator)(nil).Create), arg0, arg1)
}
