// Package pongo2gin is a template renderer that can be used with the Gin
// web framework https://github.com/gin-gonic/gin it uses the Pongo2 template
// library https://github.com/flosch/pongo2
package gin

import (
	"github.com/flosch/pongo2"
	"github.com/opensdk/gin/render"
	"net/http"
	"path"
)

// RenderOptions is used to configure the renderer.
type Pongo2RenderOptions struct {
	TemplateDir string
	ContentType string
}

// Pongo2Render is a custom Gin template renderer using Pongo2.
type Pongo2Render struct {
	Options  *Pongo2RenderOptions
	Template *pongo2.Template
	Context  pongo2.Context
}

// New creates a new Pongo2Render instance with custom Options.
func NewPongo2RenderOptions(options Pongo2RenderOptions) *Pongo2Render {
	return &Pongo2Render{
		Options: &options,
	}
}

// Default creates a Pongo2Render instance with default options.
func DefaultPongo2Render() *Pongo2Render {
	return NewPongo2RenderOptions(Pongo2RenderOptions{
		TemplateDir: "templates",
		ContentType: "text/html; charset=utf-8",
	})
}

// Instance should return a new Pongo2Render struct per request and prepare
// the template by either loading it from disk or using pongo2's cache.
func (p Pongo2Render) Instance(name string, data interface{}) render.Render {
	var template *pongo2.Template
	filename := path.Join(p.Options.TemplateDir, name)

	//	 always read template files from disk if in debug mode, use cache otherwise.
	if Mode() == "debug" {
		template = pongo2.Must(pongo2.FromFile(filename))
	} else {
		template = pongo2.Must(pongo2.FromCache(filename))
	}

	return Pongo2Render{
		Template: template,
		Context:  data.(pongo2.Context),
		Options:  p.Options,
	}
}

// Render should render the template to the response.
func (p Pongo2Render) Render(w http.ResponseWriter) error {
	writeContentType(w, []string{p.Options.ContentType})
	err := p.Template.ExecuteWriter(p.Context, w)
	return err
}

// writeContentType is also in the gin/render package but it has not been made
// pubic so is repeated here, maybe convince the author to make this public.
func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
