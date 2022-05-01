package people

import (
	"el/module"
	"html/template"
	"log"
	"net/http"
)

func PriemPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/people_priem" {
		http.NotFound(w, r)
		return
	}

	res := PriemMain()
	res.User = module.GetUserName(r)

	//inf := Info{res}
	files := []string{
		"./templates/out6.html",
		"./templates/base_layout.html",
		"./templates/footer.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
