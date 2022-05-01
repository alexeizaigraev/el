package people

import (
	"html/template"
	"log"
	"net/http"
)

func PagePeopleMenu(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/menu_people" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./templates/people/menu_people.html",
		"./templates/base_layout.html",
		"./templates/footer.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
