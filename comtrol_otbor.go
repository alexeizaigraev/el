package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func ControlOtborMenu(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/control_otbor_menu" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./templates/control_otbor_menu.html",
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

func ControlOtborTerm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/control_otbor_term" {
		http.NotFound(w, r)
		return
	}

	res := OtborTextTerm()
	res.User = GetUserName(r)

	//inf := Info{res}
	files := []string{
		"./templates/control_otbor_menu.html",
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

func OtborTextTerm() Out6 {
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
	arr, _ := FileToVec(DataInDir + "otbor_term.csv")

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
		Kind: "Отбор терминалы",
		//Data:  strings.Split(Unfind+OutText, "/n"),
		//Fname: outFname,
		//Err:   "",
		Rez: info,
	}
}

func ClearTableOtbor() error {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return err
		//panic(err)
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM otbor;")
	if err != nil {
		return err
	}
	//fmt.Println("clear table otbor")
	return nil
}
