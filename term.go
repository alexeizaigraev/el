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

type Terminal struct {
	Id               string
	Department       string
	Termial          string
	Model            string
	Serial_number    string
	Date_manufacture string
	Soft             string
	Producer         string
	Rne_rro          string
	Sealing          string
	Fiscal_number    string
	Oro_serial       string
	Oro_number       string
	Ticket_serial    string
	Ticket_1sheet    string
	Ticket_number    string
	Sending          string
	Books_arhiv      string
	Tickets_arhiv    string
	To_rro           string
	Owner_rro        string
	Register         string
	Finish           string
}

func TermCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		department := r.FormValue("department")
		termial := r.FormValue("termial")
		model := r.FormValue("model")
		serial_number := r.FormValue("serial_number")
		date_manufacture := r.FormValue("date_manufacture")
		soft := r.FormValue("soft")
		producer := r.FormValue("producer")
		rne_rro := r.FormValue("rne_rro")
		sealing := r.FormValue("sealing")
		fiscal_number := r.FormValue("fiscal_number")
		oro_serial := r.FormValue("oro_serial")
		oro_number := r.FormValue("oro_number")
		ticket_serial := r.FormValue("ticket_serial")
		ticket_1sheet := r.FormValue("ticket_1sheet")
		ticket_number := r.FormValue("ticket_number")
		sending := r.FormValue("sending")
		books_arhiv := r.FormValue("books_arhiv")
		tickets_arhiv := r.FormValue("tickets_arhiv")
		to_rro := r.FormValue("to_rro")
		owner_rro := r.FormValue("owner_rro")
		register := r.FormValue("register")
		finish := r.FormValue("finish")

		db, err := sql.Open("postgres", ConnStr)
		sqlStatement := `
INSERT INTO terminals (department, termial, model, serial_number, date_manufacture, soft, producer, rne_rro, sealing, fiscal_number, oro_serial, oro_number, ticket_serial, ticket_1sheet, ticket_number, sending, books_arhiv, tickets_arhiv, to_rro, owner_rro, register, finish)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
`

		_, err = db.Exec(sqlStatement,
			department, termial, model, serial_number, date_manufacture, soft, producer, rne_rro, sealing, fiscal_number, oro_serial, oro_number, ticket_serial, ticket_1sheet, ticket_number, sending, books_arhiv, tickets_arhiv, to_rro, owner_rro, register, finish)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/term_index", 301)
	} else {
		files := []string{
			"./templates/term/term_create.html",
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

func TermDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := sql.Open("postgres", ConnStr)
	_, err = db.Exec("delete from terminals where id=$1",
		id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/term_index", 301)
}

func TermEditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	id := r.FormValue("id")

	department := r.FormValue("department")
	termial := r.FormValue("termial")
	model := r.FormValue("model")
	serial_number := r.FormValue("serial_number")
	date_manufacture := r.FormValue("date_manufacture")
	soft := r.FormValue("soft")
	producer := r.FormValue("producer")
	rne_rro := r.FormValue("rne_rro")
	sealing := r.FormValue("sealing")
	fiscal_number := r.FormValue("fiscal_number")
	oro_serial := r.FormValue("oro_serial")
	oro_number := r.FormValue("oro_number")
	ticket_serial := r.FormValue("ticket_serial")
	ticket_1sheet := r.FormValue("ticket_1sheet")
	ticket_number := r.FormValue("ticket_number")
	sending := r.FormValue("sending")
	books_arhiv := r.FormValue("books_arhiv")
	tickets_arhiv := r.FormValue("tickets_arhiv")
	to_rro := r.FormValue("to_rro")
	owner_rro := r.FormValue("owner_rro")
	register := r.FormValue("register")
	finish := r.FormValue("finish")

	db, err := sql.Open("postgres", ConnStr)
	_, err = db.Exec("update terminals set department=$1, termial=$2, model=$3, serial_number=$4, date_manufacture=$5, soft=$6, producer=$7, rne_rro=$8, sealing=$9, fiscal_number=$10, oro_serial=$11, oro_number=$12, ticket_serial=$13, ticket_1sheet=$14, ticket_number=$15, sending=$16, books_arhiv=$17, tickets_arhiv=$18, to_rro=$19, owner_rro=$20, register=$21, finish=$22 where id = $23",
		department, termial, model, serial_number, date_manufacture, soft, producer, rne_rro, sealing, fiscal_number, oro_serial, oro_number, ticket_serial, ticket_1sheet, ticket_number, sending, books_arhiv, tickets_arhiv, to_rro, owner_rro, register, finish, id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/term_index", 301)
}

func TermEditPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	db, err := sql.Open("postgres", ConnStr)
	row := db.QueryRow("select * from terminals where id = $1", id)
	term := Terminal{}
	err = row.Scan(
		&term.Id,
		&term.Department,
		&term.Termial,
		&term.Model,
		&term.Serial_number,
		&term.Date_manufacture,
		&term.Soft,
		&term.Producer,
		&term.Rne_rro,
		&term.Sealing,
		&term.Fiscal_number,
		&term.Oro_serial,
		&term.Oro_number,
		&term.Ticket_serial,
		&term.Ticket_1sheet,
		&term.Ticket_number,
		&term.Sending,
		&term.Books_arhiv,
		&term.Tickets_arhiv,
		&term.To_rro,
		&term.Owner_rro,
		&term.Register,
		&term.Finish)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		files := []string{
			"./templates/term/term_edit.html",
			"./templates/base_layout.html",
			"./templates/footer.html",
		}

		tmpl, _ := template.ParseFiles(files...)
		err = tmpl.Execute(w, term)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

func TermIndex(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/term_index" {
		http.NotFound(w, r)
		return
	}

	db, err := sql.Open("postgres", ConnStr)
	rows, err := db.Query("select * from terminals order by termial")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	terms := []Terminal{}

	for rows.Next() {
		p := Terminal{}
		err := rows.Scan(&p.Id, &p.Department, &p.Termial, &p.Model, &p.Serial_number, &p.Date_manufacture, &p.Soft, &p.Producer, &p.Rne_rro, &p.Sealing, &p.Fiscal_number, &p.Oro_serial, &p.Oro_number, &p.Ticket_serial, &p.Ticket_1sheet, &p.Ticket_number, &p.Sending, &p.Books_arhiv, &p.Tickets_arhiv, &p.To_rro, &p.Owner_rro, &p.Register, &p.Finish)

		if err != nil {
			fmt.Println(err)
			continue
		}
		terms = append(terms, p)
	}
	files := []string{
		"./templates/term/term_index.html",
		"./templates/base_layout.html",
		"./templates/footer.html",
	}

	tmpl, _ := template.ParseFiles(files...)
	err = tmpl.Execute(w, terms)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
