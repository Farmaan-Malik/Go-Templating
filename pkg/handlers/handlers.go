package handlers

import (
	"net/http"

	"github.com/Farmaan-Malik/Go-Templating/pkg/config"
	"github.com/Farmaan-Malik/Go-Templating/pkg/models"
	"github.com/Farmaan-Malik/Go-Templating/pkg/render"
)

var Repo *Repository

type Repository struct{
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository{
	return &Repository{
		App : a,
	}
}
func NewHandler(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//business logic to generate a value -->
	//send as template data -->
	stringMap:= map[string]string{}
	stringMap["test"] = "Hello I just sent this data to the template from handler function"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
