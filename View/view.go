package View

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   string = "views/"
	TemplateExt string = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	// ... this is the same
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
func (v *View) Render(w http.ResponseWriter,
	data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}