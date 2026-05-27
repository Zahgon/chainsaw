package discovery

import (
	"regexp"
)

var StepFileName = regexp.MustCompile(`^(\d+)-(.*)\.(?:yaml|yml)$`)

type Step struct {
	AssertFiles []string
	ErrorFiles  []string
	OtherFiles  []string
}

func TryFindStepFiles(path string) (map[string]Step, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// collect and sort candidate files
