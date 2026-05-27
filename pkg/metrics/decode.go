package metrics

import (
	"github.com/prometheus/common/model"
)

func Decode(in string, ts model.Time) (model.Vector, error) {
	_ = "STUB: not implemented"
	return *new(model.Vector), nil
}
