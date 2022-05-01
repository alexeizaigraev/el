package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func OtborInputHandlerTerm(w http.ResponseWriter, r *http.Request) {
	OtborInputHandlerPapa(w, r, OtborInputTextTerm)
}

func OtborInputHandlerPapa(w http.ResponseWriter, r *http.Request,
	funcPar func(otborTextPar string) Out6) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		otborText := r.FormValue("otbor_text")
		res := funcPar(otborText)
		res.User = GetUserName(r)

		files := []string{
			"./templates/control_otbor_input_menu.html",
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

	} else {
		files := []string{
			"./templates/control_otbor_input_menu.html",
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
}

func OtborInputHandlerTerm0(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		otborText := r.FormValue("otbor_text")
		res := OtborInputTextTerm(otborText)
		res.User = GetUserName(r)

		files := []string{
			"./templates/control_otbor_input_menu.html",
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

	} else {
		files := []string{
			"./templates/control_otbor_input_menu.html",
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
}

func OtborInputTextTerm(otborText string) Out6 {
	ClearTableOtbor()
	info := ""
	count := 0
	count_err := 0
	ok := true
	term := ""
	dep := ""
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var arr []string
	if strings.Contains(otborText, "\n") {
		arr = strings.Split(otborText, "\n")
	} else {
		arr = append(arr, strings.Trim(otborText, " "))
	}

	for _, term0 := range arr {

		if term0 == "" || term0 == "\n" {
			continue
		}
		term = strings.Trim(term0, " ")
		term = strings.Trim(term, "\n")
		dep = string([]rune(term)[:7])

		sqlStatement := `
INSERT INTO otbor (term, dep)
VALUES ($1, $2)
RETURNING term`

		myTerm := ""
		err = db.QueryRow(sqlStatement, term, dep).Scan(&myTerm)
		if err != nil {
			fmt.Println(err)
			ok = false
			count_err += 1
		}
		count += 1
		fmt.Println(term, dep)
	}

	if ok {
		info += fmt.Sprintf("success otbor %d", count)
		//fmt.Println("success")
	} else {
		info += fmt.Sprintf(">> otbor err: %d, ok: %d", count_err, count)
		//fmt.Println("errors")
	}
	return Out6{
		Kind: "Отбор ввод терминалы",
		//Data:  strings.Split(Unfind+OutText, "/n"),
		//Fname: outFname,
		//Err:   "",
		Rez: info,
	}
}
