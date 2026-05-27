package model

import (
	"sync"
	"time"
)

type OperationType string

const (
	OperationTypeApply   OperationType = "apply"
	OperationTypeAssert  OperationType = "assert"
	OperationTypeCommand OperationType = "command"
	OperationTypeCreate  OperationType = "create"
	OperationTypeDelete  OperationType = "delete"
	OperationTypeError   OperationType = "error"
	OperationTypePatch   OperationType = "patch"
	OperationTypeScript  OperationType = "script"
	OperationTypeSleep   OperationType = "sleep"
	OperationTypeUpdate  OperationType = "update"
)

type Report struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Tests     []*TestReport
	lock      sync.Mutex
}

func (r *Report) Add(report *TestReport) { _ = "STUB: not implemented"; return }

type TestReport struct {
	BasePath   string
	Name       string
	Concurrent *bool
	StartTime  time.Time
	EndTime    time.Time
	Namespace  string
	Skipped    bool
	Steps      []*StepReport
}

func (r *TestReport) Add(report *StepReport) { _ = "STUB: not implemented"; return }

type StepReport struct {
	Name       string
	StartTime  time.Time
	EndTime    time.Time
	Operations []*OperationReport
}

func (r *StepReport) Add(report *OperationReport) { _ = "STUB: not implemented"; return }

func (r *StepReport) Failed() bool { _ = "STUB: not implemented"; return false }

func (r *TestReport) Failed() bool { _ = "STUB: not implemented"; return false }

type OperationReport struct {
	Name      string
	Type      OperationType
	StartTime time.Time
	EndTime   time.Time
	Err       error
}
