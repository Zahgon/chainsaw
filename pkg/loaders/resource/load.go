package resource

import (
	"io"
	"net/url"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type (
	splitter   = func([]byte) ([][]byte, error)
	converter  = func([]byte) ([]byte, error)
	readSeeker = interface {
		io.Reader
		io.Seeker
	}
)

func readDownloadedContent(reader readSeeker) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Load(pattern string, manifest bool) ([]unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadFromURI(url *url.URL, manifest bool) ([]unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Parse(content []byte, manifest bool) ([]unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parse(content []byte, splitter splitter, converter converter, manifest bool) ([]unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
