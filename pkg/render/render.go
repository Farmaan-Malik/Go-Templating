package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Farmaan-Malik/Go-Templating/pkg/config"
)
var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app = a
}
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template

	if app.UseCache{
		tc= app.TemplateCache
	}else{
		tc,_=CreateTemplateCache()
	}
	
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("couldn't get template from template cache")
	}
	//create a buffer
	buf := new(bytes.Buffer)
	//execute the template file in t and store the data in a buffer "buf"
	_ = t.Execute(buf, nil)
	//render the template by writing the data in the buffer to the esponseWriter
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser")
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all the files having .page.tmpl in their name and save that to slice "pages"
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range over pages and get the name of the file without the full path-- to use as a key in our map along with the parsed files as their value
	for _, page := range pages {
		//getting the name from the files
		name := filepath.Base(page)
		//making a new template with the name of the file and its parsed data
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		//checking if a layout file exists
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		// layout exists
		if len(matches) > 0 {
			//add the layout file to the ts variable which now has parsed page and parsed layout
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		//add the file to the map "myCache" with the base name of the page template as key and parsed page & layout as value
		myCache[name] = ts
	}
	return myCache, nil

}
