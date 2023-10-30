package files

import (
	"html/template"
	"net/http"
)

type paths struct {
	baseName string
	pagePath string
	basePath string
}

type Files struct{}

func getPagePath(template string) *paths {
	p := &paths{}
	p.basePath = baseTemplate
	p.pagePath = template
	p.baseName = base
	return p
}

func (f *Files) RenderFile(w http.ResponseWriter, path string) error {
	paths := getPagePath(path)
	files := []string{
		paths.basePath,
		paths.pagePath,
	}

	pg, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	err = pg.ExecuteTemplate(w, paths.baseName, nil)
	return err
}
