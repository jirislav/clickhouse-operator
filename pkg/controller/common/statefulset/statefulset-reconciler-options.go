// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
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

package statefulset

type ReconcileStatefulSetOptions struct {
	forceRecreate bool
}

func (o *ReconcileStatefulSetOptions) ForceRecreate() bool {
	if o == nil {
		return false
	}
	return o.forceRecreate
}

func NewReconcileStatefulSetOptions(forceRecreate bool) *ReconcileStatefulSetOptions {
	return &ReconcileStatefulSetOptions{
		forceRecreate: forceRecreate,
	}
}

type ReconcileStatefulSetOptionsArr []*ReconcileStatefulSetOptions

// NewReconcileStatefulSetOptionsArr creates new reconcileHostStatefulSetOptions array
func NewReconcileStatefulSetOptionsArr(opts ...*ReconcileStatefulSetOptions) (res ReconcileStatefulSetOptionsArr) {
	return append(res, opts...)
}

// First gets first option
func (a ReconcileStatefulSetOptionsArr) First() *ReconcileStatefulSetOptions {
	if len(a) > 0 {
		return a[0]
	}
	return nil
}