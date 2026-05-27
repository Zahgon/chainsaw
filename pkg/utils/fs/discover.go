package fs

import (
	"os"
	"path/filepath"
)

func discoverFolders(stat func(string) (os.FileInfo, error), walk func(string, filepath.WalkFunc) error, paths ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DiscoverFolders(paths ...string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
