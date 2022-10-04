package html

import (
	"embed"
	"io"
	"text/template"
)

//go:embed *
var files embed.FS

var (
	chat = parse("chat.html")
)

type ChatParams struct {
	Title string
}

func Chat(w io.Writer, p ChatParams) error {
	return chat.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}
