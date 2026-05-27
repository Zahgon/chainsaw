package internal

import (
	"io"
	"reflect"
	"regexp"
	"time"
)

// TestDeps implements the TestDeps interface for MainStart.
type TestDeps struct {
	Test     bool
	matchPat string
	matchRe  *regexp.Regexp
}

func (d *TestDeps) MatchString(pat, str string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (*TestDeps) SetPanicOnExit0(bool) { _ = "STUB: not implemented"; return }

func (*TestDeps) StartCPUProfile(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (*TestDeps) StopCPUProfile() { _ = "STUB: not implemented"; return }

func (*TestDeps) WriteProfileTo(name string, w io.Writer, debug int) error {
	_ = "STUB: not implemented"
	return nil
}

func (*TestDeps) ImportPath() string { _ = "STUB: not implemented"; return "" }

func (*TestDeps) ModulePath() string { _ = "STUB: not implemented"; return "" }

func (*TestDeps) StartTestLog(w io.Writer) { _ = "STUB: not implemented"; return }

func (*TestDeps) StopTestLog() error { _ = "STUB: not implemented"; return nil }

func (*TestDeps) CoordinateFuzzing(time.Duration, int64, time.Duration, int64, int, []corpusEntry, []reflect.Type, string, string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*TestDeps) RunFuzzWorker(func(corpusEntry) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (*TestDeps) ReadCorpus(string, []reflect.Type) ([]corpusEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*TestDeps) CheckCorpus([]any, []reflect.Type) error { _ = "STUB: not implemented"; return nil }

func (*TestDeps) ResetCoverage() { _ = "STUB: not implemented"; return }

func (*TestDeps) SnapshotCoverage() { _ = "STUB: not implemented"; return }

func (*TestDeps) InitRuntimeCoverage() (string, func(string, string) (string, error), func() float64) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
