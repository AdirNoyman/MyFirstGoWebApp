package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Preparation for custom functions I might insert to the templates
var functions = template.FuncMap{}

// RenderTemplate Render the html templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()

	if err != nil {
		// If there is an error, kill the app because I can't show anything to the user
		log.Fatal("There is an error. Killing the app 😫", err)
	}

	t, ok := tc[tmpl]

	// Checking if the name of the template that was passed is in the template cache
	if !ok {

		log.Fatal("There is an error. Killing the app 😫", err)
	}

	// Turn the parsed template we found in to bytes
	buf := new(bytes.Buffer)

	// Storing the bytes we got in to this buf variable
	_ = t.Execute(buf, nil)

	// Writing the template into the browser
	_, err = buf.WriteTo(w)

	if err != nil {

		fmt.Println("Error writing template to browser", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)

	if err != nil {

		fmt.Println("Error parsing template 🥴:", err)
		return
	}

}

// CreateTemplateCache creates templates cache as a map(name of template : parsed template)
func CreateTemplateCache() (map[string]*template.Template, error) {

	// This creates a map where the string is the name(key) of the template and the value will the parsed template
	myCache := map[string]*template.Template{}

	// Get all the files in the templates folder that start with the word page
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {

		return myCache, err
	}

	// _ = index, page = template name
	for _, page := range pages {

		name := filepath.Base(page)

		// ts = template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {

			return myCache, err
		}

		// Look in the templates for any file that contains layout
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {

			return myCache, err
		}

		if len(matches) > 0 {

			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}

		myCache[name] = ts
	}

	return myCache, nil

}
