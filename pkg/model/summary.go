package model

import (
	"sync/atomic"
)

type SummaryResult interface {
	Passed() int32
	Failed() int32
	Skipped() int32
}

type Summary struct {
	passed  atomic.Int32
	failed  atomic.Int32
	skipped atomic.Int32
}

func (s *Summary) IncPassed() { _ = "STUB: not implemented"; return }

func (s *Summary) IncFailed() { _ = "STUB: not implemented"; return }

func (s *Summary) IncSkipped() { _ = "STUB: not implemented"; return }

func (s *Summary) Passed() int32 { _ = "STUB: not implemented"; return 0 }

func (s *Summary) Failed() int32 { _ = "STUB: not implemented"; return 0 }

func (s *Summary) Skipped() int32 { _ = "STUB: not implemented"; return 0 }
