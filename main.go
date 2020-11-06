package main

import (
	"fmt"
	"github.com/wyster/go-sendgrid/template"
	"log"
	"net/http"
	os "os"
)

func templatesListHandler(w http.ResponseWriter, r *http.Request) {
	responseData := template.List(os.Getenv("SENDGRID_TOKEN"))
	fmt.Fprintln(w, "<html><ul>")
	for _, id := range responseData.IdsList() {
		fmt.Fprintf(w, "<li><a href='/template/show/%s'>%s</a></li>", id, id)
	}
	fmt.Fprintln(w, "</ul></html>")
}

func templateShowHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/template/show/"):]
	responseData := template.Get(os.Getenv("SENDGRID_TOKEN"), id)
	if len(responseData.Versions) > 0 && responseData.Versions[0].Active == 1 {
		fmt.Fprintf(w, "%s", responseData.Versions[0].HtmlContent)
	} else {
		fmt.Fprintf(w, "%s", "Active template version not found")
	}
}

func main() {
	fmt.Println("started")
	http.HandleFunc("/template/list/", templatesListHandler)
	http.HandleFunc("/template/show/", templateShowHandler)
	httpPort := "80"
	if len(os.Getenv("HTTP_PORT")) > 0 {
		httpPort = os.Getenv("HTTP_PORT")
	}
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
