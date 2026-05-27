package report

import (
	"time"

	"github.com/kyverno/chainsaw/pkg/model"
)

func durationInSecondsString(start, end time.Time) string { _ = "STUB: not implemented"; return "" }

func saveJUnitTest(report *model.Report, file string) error { _ = "STUB: not implemented"; return nil }

func saveJUnitStep(report *model.Report, file string) error { _ = "STUB: not implemented"; return nil }

func saveJUnitOperation(report *model.Report, file string) error {
	_ = "STUB: not implemented"
	return nil
}
