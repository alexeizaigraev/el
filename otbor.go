// electrum project main.go
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Otbor struct {
	Id   int
	Term string
	Dep  string
}

func OtborDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := sql.Open("postgres", ConnStr)
	_, err = db.Exec("delete from otbor where id=$1",
		id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/otbor_index", 301)
}

func OtborEditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	id := r.FormValue("id")
	term := r.FormValue("term")
	dep := r.FormValue("dep")

	db, err := sql.Open("postgres", ConnStr)
	_, err = db.Exec("update otbor set term=$1, dep=$2 where id = $3",
		term, dep, id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/otbor_index", 301)
}

func OtborEditPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	db, err := sql.Open("postgres", ConnStr)
	row := db.QueryRow("select * from otbor where id = $1", id)
	otbor := Otbor{}
	err = row.Scan(&otbor.Id, &otbor.Term, &otbor.Dep)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		files := []string{
			"./templates/otbor/otbor_edit.html",
			"./templates/base_layout.html",
			"./templates/footer.html",
		}

		tmpl, _ := template.ParseFiles(files...)
		err = tmpl.Execute(w, otbor)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

func OtborIndex(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/otbor_index" {
		http.NotFound(w, r)
		return
	}

	db, err := sql.Open("postgres", ConnStr)
	rows, err := db.Query("select * from otbor order by term")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	otbors := []Otbor{}

	for rows.Next() {
		p := Otbor{}
		err := rows.Scan(&p.Id, &p.Term, &p.Dep)
		if err != nil {
			fmt.Println(err)
			continue
		}
		otbors = append(otbors, p)
	}
	files := []string{
		"./templates/otbor/otbor_index.html",
		"./templates/base_layout.html",
		"./templates/footer.html",
	}

	tmpl, _ := template.ParseFiles(files...)
	err = tmpl.Execute(w, otbors)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
