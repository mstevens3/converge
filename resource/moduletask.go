// Copyright © 2016 Asteris, LLC
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

package resource

// ModuleTask is the task for calling a module.
type ModuleTask struct {
	Args       Values
	Source     string
	ModuleName string

	parent *Module
}

// Name returns name for metadata
func (m *ModuleTask) Name() string {
	return m.ModuleName
}

// Validate checks shell tasks validity
func (m *ModuleTask) Validate() error {
	return nil
}

// Prepare this module for use
func (m *ModuleTask) Prepare(parent *Module) error {
	m.parent = parent
	return nil
}
