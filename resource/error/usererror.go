// Copyright © 2017 Asteris, LLC
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

package usererror

import (
	"github.com/asteris-llc/converge/load/registry"
	"github.com/asteris-llc/converge/resource"
	"golang.org/x/net/context"
)

// Preparer for Error
//
// Generates a runtime error in the graph
type Preparer struct {
	Error string `hcl:"error" required:"true"`
}

// ApplyPreparer for error.apply
//
// Generates a runtime error in the graph
type ApplyPreparer struct {
	Error string `hcl:"error" required:"true"`
}

// Prepare creates a new UserError
func (p *Preparer) Prepare(context.Context, resource.Renderer) (resource.Task, error) {
	return &UserError{Error: p.Error}, nil
}

// Prepare creates a new UserError
func (p *ApplyPreparer) Prepare(context.Context, resource.Renderer) (resource.Task, error) {
	return &UserError{Error: p.Error, SkipPlan: true}, nil
}

// UserError implements resource.Task for a user error
type UserError struct {
	Error    string
	SkipPlan bool
	changed  bool
}

// Check returns an error during the plan phase
func (u *UserError) Check(context.Context, resource.Renderer) (resource.TaskStatus, error) {
	status := resource.NewStatus()
	if u.SkipPlan {
		if !u.changed {
			status.RaiseLevel(resource.StatusWillChange)
		}
		return status, nil
	}
	status.RaiseLevel(resource.StatusFatal)
	status.AddMessage("runtime error: explicit error encountered")
	status.AddMessage(u.Error)
	return status, nil
}

// Apply returns an error during the apply phase
func (u *UserError) Apply(context.Context) (resource.TaskStatus, error) {
	u.changed = true
	status := resource.NewStatus()
	status.RaiseLevel(resource.StatusFatal)
	status.AddMessage("runtime error: explicit error encountered")
	status.AddMessage(u.Error)
	return status, nil
}

func init() {
	registry.Register("error", (*Preparer)(nil), (*UserError)(nil))
	registry.Register("error.plan", (*Preparer)(nil), (*UserError)(nil))
	registry.Register("error.apply", (*ApplyPreparer)(nil), (*UserError)(nil))
}