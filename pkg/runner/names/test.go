package names

import (
	"github.com/kyverno/chainsaw/pkg/discovery"
)

type (
	workignDirInterface   = func() (string, error)
	absolutePathInterface = func(string) (string, error)
	relativePathInterface = func(string, string) (string, error)
)

func Test(test discovery.Test, full bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func helpTest(test discovery.Test, workingDir workignDirInterface, absolutePath absolutePathInterface, relativePath relativePathInterface) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
