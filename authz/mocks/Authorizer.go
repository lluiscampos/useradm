// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package mocks

import authz "github.com/mendersoftware/useradm/authz"
import context "context"
import jwt "github.com/mendersoftware/useradm/jwt"
import mock "github.com/stretchr/testify/mock"

// Authorizer is an autogenerated mock type for the Authorizer type
type Authorizer struct {
	mock.Mock
}

// Authorize provides a mock function with given fields: ctx, token, resource, action
func (_m *Authorizer) Authorize(ctx context.Context, token *jwt.Token, resource string, action string) error {
	ret := _m.Called(ctx, token, resource, action)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *jwt.Token, string, string) error); ok {
		r0 = rf(ctx, token, resource, action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ authz.Authorizer = (*Authorizer)(nil)
