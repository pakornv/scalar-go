package scalar

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	DefaultTitle = "Scalar API Reference"
	DefaultCDN   = "https://cdn.jsdelivr.net/npm/@scalar/api-reference"
)

type Scalar struct {
	title     string
	cdn       string
	customCSS string
	config    string
	content   string
}

// New creates a new instance of the Scalar API reference
func New(specURL string, cfg *Config) (*Scalar, error) {
	if specURL == "" {
		return nil, fmt.Errorf("spec URL or content is required")
	}

	type Spec struct {
		URL string `json:"url"`
	}

	var spec *Spec
	var content string
	if strings.HasPrefix(specURL, "http") {
		spec = &Spec{URL: specURL}
	} else {
		data, err := os.ReadFile(specURL)
		if err != nil {
			return nil, err
		}

		content = string(data)
	}

	title := cfg.Title
	if title == "" {
		title = DefaultTitle
	}

	cdn := cfg.CDN
	if cdn == "" {
		cdn = DefaultCDN
	}

	customCSS := cfg.CustomCSS
	if cfg.Theme != "" {
		customCSS = ""
	}

	// Serialize the configuration
	data, err := json.Marshal(struct {
		*Config
		Spec *Spec `json:"spec,omitempty"`
	}{
		Config: cfg,
		Spec:   spec,
	})
	if err != nil {
		return nil, err
	}

	// Replace double quotes with HTML entity
	config := strings.ReplaceAll(string(data), `"`, `&quot;`)

	return &Scalar{
		title:     title,
		cdn:       cdn,
		customCSS: customCSS,
		config:    config,
		content:   content,
	}, nil
}

// RenderHTML renders the HTML content for the API reference
func (s *Scalar) RenderHTML() (string, error) {
	htmlContent := fmt.Sprintf(`
	  <!DOCTYPE html>
	  <html lang="en">
	    <head>
	      <title>%s</title>
	      <meta charset="utf-8" />
	      <meta name="viewport" content="width=device-width, initial-scale=1" />
	      <style>%s</style>
	    </head>
	    <body>
	      <script id="api-reference" type="application/json" data-configuration="%s">%s</script>
	      <script src="%s"></script>
	    </body>
	  </html>
	`, s.title, s.customCSS, s.config, s.content, s.cdn,
	)

	return htmlContent, nil
}
