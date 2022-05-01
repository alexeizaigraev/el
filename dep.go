// electrum project main.go
package main

import (
	"database/sql"

	//"electrum/db"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Department struct {
	Id              string
	Department      string
	Region          string
	District_region string
	District_city   string
	City_type       string
	City            string
	Street          string
	Street_type     string
	Hous            string
	Post_index      string
	Partner         string
	Status          string
	Register        string
	Edrpou          string
	Address         string
	Partner_name    string
	Id_terminal     string
	Koatu           string
	Tax_id          string
	Koatu2          string
}

func DepCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		department := r.FormValue("department")
		region := r.FormValue("region")
		district_region := r.FormValue("district_region")
		district_city := r.FormValue("district_city")

		city_type := r.FormValue("city_type")
		city := r.FormValue("city")
		street := r.FormValue("street")

		street_type := r.FormValue("street_type")
		hous := r.FormValue("hous")
		post_index := r.FormValue("post_index")

		partner := r.FormValue("partner")

		status := r.FormValue("status")
		register := r.FormValue("register")
		edrpou := r.FormValue("edrpou")
		address := r.FormValue("address")
		partner_name := r.FormValue("partner_name")
		id_terminal := r.FormValue("id_terminal")
		koatu := r.FormValue("koatu")
		tax_id := r.FormValue("tax_id")
		koatu2 := r.FormValue("koatu2")

		db, err := sql.Open("postgres", ConnStr)
		sqlStatement := `
INSERT INTO departments (department, region, district_region, district_city, city_type, city, street, street_type, hous, post_index, partner, status, register, edrpou, address, partner_name, id_terminal, koatu, tax_id, koatu2)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
`

		_, err = db.Exec(sqlStatement,
			department, region, district_region, district_city, city_type, city, street, street_type, hous, post_index, partner, status, register, edrpou, address, partner_name, id_terminal, koatu, tax_id, koatu2)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/dep_index", 301)
	} else {
		files := []string{
			"./templates/dep/dep_create.html",
			"./templates/base_layout.html",
			"./templates/footer.html",
		}

		tmpl, _ := template.ParseFiles(files...)
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}

}

func DepDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := sql.Open("postgres", ConnStr)
	_, err = db.Exec("delete from departments where id=$1",
		id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/dep_index", 301)
}

func DepEditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	id := r.FormValue("id")

	department := r.FormValue("department")
	region := r.FormValue("region")
	district_region := r.FormValue("district_region")
	district_city := r.FormValue("district_city")
	city_type := r.FormValue("city_type")
	city := r.FormValue("city")
	street := r.FormValue("street")
	street_type := r.FormValue("street_type")
	hous := r.FormValue("hous")
	post_index := r.FormValue("post_index")
	partner := r.FormValue("partner")
	status := r.FormValue("status")
	register := r.FormValue("register")
	edrpou := r.FormValue("edrpou")
	address := r.FormValue("address")
	partner_name := r.FormValue("partner_name")
	id_terminal := r.FormValue("id_terminal")
	koatu := r.FormValue("koatu")
	tax_id := r.FormValue("tax_id")
	koatu2 := r.FormValue("koatu2")

	db, err := sql.Open("postgres", ConnStr)
	_, err = db.Exec("update departments set department=$1, region=$2, district_region=$3, district_city=$4, city_type=$5, city=$6, street=$7, street_type=$8, hous=$9, post_index=$10, partner=$11, status=$12, register=$13, edrpou=$14, address=$15, partner_name=$16, id_terminal=$17, koatu=$18, tax_id=$19, koatu2=$20 where id = $21",
		department, region, district_region, district_city, city_type, city, street, street_type, hous, post_index, partner, status, register, edrpou, address, partner_name, id_terminal, koatu, tax_id, koatu2, id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/dep_index", 301)
}

func DepEditPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	db, err := sql.Open("postgres", ConnStr)
	row := db.QueryRow("select * from departments where id = $1", id)
	dep := Department{}
	err = row.Scan(&dep.Id, &dep.Department, &dep.Region, &dep.District_region, &dep.District_city, &dep.City_type, &dep.City, &dep.Street, &dep.Street_type, &dep.Hous, &dep.Post_index, &dep.Partner, &dep.Status, &dep.Register, &dep.Edrpou, &dep.Address, &dep.Partner_name, &dep.Id_terminal, &dep.Koatu, &dep.Tax_id, &dep.Koatu2)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		files := []string{
			"./templates/dep/dep_edit.html",
			"./templates/base_layout.html",
			"./templates/footer.html",
		}

		tmpl, _ := template.ParseFiles(files...)
		err = tmpl.Execute(w, dep)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

func DepIndex(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/dep_index" {
		http.NotFound(w, r)
		return
	}

	db, err := sql.Open("postgres", ConnStr)
	rows, err := db.Query("select * from departments order by department")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	deps := []Department{}

	for rows.Next() {
		p := Department{}
		err := rows.Scan(&p.Id, &p.Department, &p.Region, &p.District_region, &p.District_city, &p.City_type, &p.City, &p.Street, &p.Street_type, &p.Hous, &p.Post_index, &p.Partner, &p.Status, &p.Register, &p.Edrpou, &p.Address, &p.Partner_name, &p.Id_terminal, &p.Koatu, &p.Tax_id, &p.Koatu2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		deps = append(deps, p)
	}
	files := []string{
		"./templates/dep/dep_index.html",
		"./templates/base_layout.html",
		"./templates/footer.html",
	}

	tmpl, _ := template.ParseFiles(files...)
	err = tmpl.Execute(w, deps)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
