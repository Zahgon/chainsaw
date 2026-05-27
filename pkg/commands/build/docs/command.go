package docs

import (
	_ "embed"
	"io"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/kyverno/chainsaw/pkg/discovery"
	"github.com/spf13/cobra"
)

//go:embed docs.tmpl
var docsTemplate string

//go:embed catalog.tmpl
var catalogTemplate string

var (
	funcMap     = map[string]any{"fpRel": filepath.Rel, "fpJoin": filepath.Join}
	docsTmpl    = template.Must(template.New("docs").Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(docsTemplate))
	catalogTmpl = template.Must(template.New("catalog").Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(catalogTemplate))
)

type options struct {
	testFile   string
	readmeFile string
	catalog    string
	testDirs   []string
}

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func generateDocs(out io.Writer, fileName string, tests map[string][]discovery.Test) error {
	_ = "STUB: not implemented"
	return nil
}

func generateCatalog(_ io.Writer, readme string, catalog string, tests ...discovery.Test) error {
	_ = "STUB: not implemented"
	return nil
}
