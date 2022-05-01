// electrum project main.go
package main

import (
	"database/sql"
	"el/people"

	//"electrum/people"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Db *sql.DB

func main() {

	Db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		//return err
		panic(err)
	}
	defer Db.Close()

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/login", loginPage)
	router.HandleFunc("/logout", logoutPage)

	router.HandleFunc("/menu_people", people.PagePeopleMenu)
	router.HandleFunc("/people_priem", people.PriemPage)
	router.HandleFunc("/people_otpusk", people.OtpuskPage)
	router.HandleFunc("/people_perevod", people.PerevodPage)

	router.HandleFunc("/control_otbor_menu", ControlOtborMenu)
	router.HandleFunc("/control_otbor_term", ControlOtborTerm)

	router.HandleFunc("/control_otbor_input_term", OtborInputHandlerTerm)

	router.HandleFunc("/database", PageDatabase)
	router.HandleFunc("/otbor_refresh", OtborRefresh)

	router.HandleFunc("/otbor_index", OtborIndex)
	router.HandleFunc("/otbor_edit/{id:[0-9]+}", OtborEditPage).Methods("GET")
	router.HandleFunc("/otbor_edit/{id:[0-9]+}", OtborEditHandler).Methods("POST")
	router.HandleFunc("/otbor_delete/{id:[0-9]+}", OtborDeleteHandler)

	router.HandleFunc("/dep_index", DepIndex)
	router.HandleFunc("/dep_create", DepCreateHandler)
	router.HandleFunc("/dep_edit/{id:[0-9]+}", DepEditPage).Methods("GET")
	router.HandleFunc("/dep_edit/{id:[0-9]+}", DepEditHandler).Methods("POST")
	router.HandleFunc("/dep_delete/{id:[0-9]+}", DepDeleteHandler)

	router.HandleFunc("/term_index", TermIndex)
	router.HandleFunc("/term_create", TermCreateHandler)
	router.HandleFunc("/term_edit/{id:[0-9]+}", TermEditPage).Methods("GET")
	router.HandleFunc("/term_edit/{id:[0-9]+}", TermEditHandler).Methods("POST")
	router.HandleFunc("/term_delete/{id:[0-9]+}", TermDeleteHandler)

	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8000", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./templates/home_page.html",
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

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println(username, password)

		expiration := time.Now().Add(10 * time.Hour)
		cookie := http.Cookie{
			Name:    "session_id",
			Value:   username,
			Expires: expiration,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.ServeFile(w, r, "templates/login_simple.html")
	}
}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	http.Redirect(w, r, "/", http.StatusFound)
}

func OtborRefresh(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/otbor_refresh" {
		http.NotFound(w, r)
		return
	}

	ClearTableOtbor()
	res := InsertOtborFromFile() + "  |  "
	ClearTableTerminals()
	res += InsertTerminalsFromFile() + "  |  "
	ClearTableDepartments()
	res += InsertDepartmentsFromFile()

	//inf := Info{res}
	files := []string{
		"./templates/otbor/otbor_refresh.html",
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

func PageDatabase(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/database" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./templates/otbor/menu_db.html",
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
