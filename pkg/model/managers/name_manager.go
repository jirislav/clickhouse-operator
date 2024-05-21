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

package managers

import "github.com/altinity/clickhouse-operator/pkg/model/chi/namer"

type INameManager interface {
	Names(what namer.NameType, params ...any) []string
	Name(what namer.NameType, params ...any) string
}

type NameManagerType string

const (
	NameManagerTypeClickHouse NameManagerType = "clickhouse"
	NameManagerTypeKeeper     NameManagerType = "keeper"
)

func NewNameManager(what NameManagerType) INameManager {
	switch what {
	case NameManagerTypeClickHouse:
		return namer.NewClickHouse()
	}
	panic("unknown name manager type")
}