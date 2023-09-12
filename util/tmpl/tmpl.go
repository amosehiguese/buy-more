package tmpl

import (
	"fmt"
	"html/template"
	"net/http"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) error {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w,"layout",data)
	if err != nil {
		return err
	}
	return nil
}
