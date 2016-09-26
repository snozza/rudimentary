package renderer

import (
  "github.com/snozza/rudimentary/domain"
  "github.com/unrolled/render"
  "net/http"
)

const RendererKey domain.ContextKey = "rudimentary-mddlwr-unrolled-render-key"
const JSON = "json"
const Text = "text"

type Options render.Options

// Renderer type
// implements IRender and IContextMiddleware
type Renderer struct {
  r *render.Render
  options *Options
  DefaultRenderType string
}

func New(options *Options, defaultRenderType string) *Renderer {
  r := render.New(render.Options(*options))
  return &Renderer{r, options, defaultRenderType}
}

func (renderer *Renderer) Render(w http.ResponseWriter, req *http.Request, status int, v interface{}) {
  acceptHeaders := domain.NewAcceptHeadersFromString(req.Header.Get("accept"))

  renderType := renderer.DefaultRenderType
  for _, h := range acceptHeaders {
    m := h.MediaType
    if m.SubType == JSON || m.Suffix == JSON {
      renderType = JSON
      break
    }
  }

  switch renderType {
  case JSON:
    renderer.JSON(w, status, v)
  default:
    renderer.Text(w, status, v.([]byte))
  }
}

func (renderer *Renderer) JSON(w http.ResponseWriter, status int, v interface{}) {
  renderer.r.JSON(w, status, v)
}

func (renderer *Renderer) Text(w http.ResponseWriter, status int, v []byte) {
  w.WriteHeader(status)
  w.Write(v)
}
