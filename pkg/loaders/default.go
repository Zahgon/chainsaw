package loaders

import (
	"io/fs"
	"sync"

	"github.com/kyverno/pkg/ext/resource/loader"
)

func defaultLoader(_fs func() (fs.FS, error)) (loader.Loader, error) {
	_ = "STUB: not implemented"
	return *new(loader.Loader), nil
}

var DefaultLoader = sync.OnceValues(func() (loader.Loader, error) { return defaultLoader(nil) })
