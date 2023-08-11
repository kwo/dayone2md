package dayone2md

import (
	"embed"
	"fmt"
	"log/slog"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

func loadTemplate(name string) (*template.Template, error) {
	if strings.Contains(name, ".") {
		slog.Debug("loading external template", "template", name)
		tpl, err := template.ParseFiles(name)
		if err != nil {
			return nil, fmt.Errorf("cannot load external template (%s): %w", name, err)
		}
		return tpl, nil
	}

	slog.Debug("loading builtin template", "template", name)
	templateName := fmt.Sprintf("templates/%s.tmpl", name)
	tpl, err := template.ParseFS(templatesFS, templateName)
	if err != nil {
		return nil, fmt.Errorf("cannot load builtin template (%s): %w", name, err)
	}
	return tpl, nil
}
