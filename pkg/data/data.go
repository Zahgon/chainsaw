package data

import (
	"embed"
	"io/fs"
	"sync"
)

const (
	configFile    = "default.yaml"
	configFolder  = "config"
	crdsFolder    = "crds"
	schemasFolder = "schemas/json"
)

//go:embed crds
var crdsFs embed.FS

//go:embed config
var configFs embed.FS

//go:embed schemas
var schemasFs embed.FS

func _config() (fs.FS, error) { _ = "STUB: not implemented"; return *new(fs.FS), nil }

func _configFile(_fs func() (fs.FS, error)) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _crds() (fs.FS, error) { _ = "STUB: not implemented"; return *new(fs.FS), nil }

func _schemas() (fs.FS, error) { _ = "STUB: not implemented"; return *new(fs.FS), nil }

func _sub(f embed.FS, dir string) (fs.FS, error) {
	_ = "STUB: not implemented"
	return *new(fs.FS), nil
}

var (
	config     = sync.OnceValues(_config)
	ConfigFile = sync.OnceValues(func() ([]byte, error) { return _configFile(nil) })
	Crds       = sync.OnceValues(_crds)
	Schemas    = sync.OnceValues(_schemas)
)
