package handler

import (
	"app/model"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
)

var context *model.TemplateContext = &model.TemplateContext{
	Error: "",
	Data:  nil,
}

func clearErrorMessage() {
	context.Error = ""
}

func renderTemplate(w http.ResponseWriter, templatePath string) {
	t, _ := template.ParseFiles("template/base.html", templatePath)
	err := t.Execute(w, context)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readableFileSize(size int64) string {
	i := math.Floor(math.Log(float64(size)) / math.Log(1024))
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	return fmt.Sprintf("%.02F %s", float64(size)/math.Pow(1024, i), sizes[int(i)])
}

func redirectWithErrorMessage(message string, w http.ResponseWriter, r *http.Request) {
	context.Error = message
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
