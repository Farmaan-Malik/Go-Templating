package handlers

import (
	"net/http"
	"github.com/Farmaan-Malik/Go-Templating/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"home.page.html")
}
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"about.page.html")
}
