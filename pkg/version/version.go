package version

import (
	"runtime/debug"
)

const (
	notFound    = "---"
	vcsTime     = "vcs.time"
	vcsRevision = "vcs.revision"
)

// BuildVersion is provided at compile-time
var BuildVersion string

type buildInfoReader = func() (*debug.BuildInfo, bool)

func Version() string { _ = "STUB: not implemented"; return "" }

func version(reader buildInfoReader) string { _ = "STUB: not implemented"; return "" }

func Time() string { _ = "STUB: not implemented"; return "" }

func time(reader buildInfoReader) string { _ = "STUB: not implemented"; return "" }

func Hash() string { _ = "STUB: not implemented"; return "" }

func hash(reader buildInfoReader) string { _ = "STUB: not implemented"; return "" }

func tryFindSetting(key string, settings ...debug.BuildSetting) string {
	_ = "STUB: not implemented"
	return ""
}
