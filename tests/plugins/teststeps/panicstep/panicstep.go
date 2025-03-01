// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package panicstep

import (
	"encoding/json"

	"github.com/linuxboot/contest/pkg/event"
	"github.com/linuxboot/contest/pkg/event/testevent"
	"github.com/linuxboot/contest/pkg/test"
	"github.com/linuxboot/contest/pkg/xcontext"
)

// Name is the name used to look this plugin up.
var Name = "Panic"

// Events defines the events that a TestStep is allow to emit
var Events = []event.Name{}

type panicStep struct {
}

// Name returns the name of the Step
func (ts *panicStep) Name() string {
	return Name
}

// Run executes the example step.
func (ts *panicStep) Run(
	ctx xcontext.Context,
	ch test.TestStepChannels,
	ev testevent.Emitter,
	stepsVars test.StepsVariables,
	inputParams test.TestStepParameters,
	resumeState json.RawMessage,
) (json.RawMessage, error) {
	panic("panic step")
}

// ValidateParameters validates the parameters associated to the TestStep
func (ts *panicStep) ValidateParameters(_ xcontext.Context, params test.TestStepParameters) error {
	return nil
}

// New creates a new panicStep
func New() test.TestStep {
	return &panicStep{}
}
