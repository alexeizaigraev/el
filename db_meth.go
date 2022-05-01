package main

import (
	"database/sql"
	"fmt"
)

func InsertDepartmentsFromFile() string {
	info := ""
	count := 0
	countErr := 0
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println(err)
		//return err
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO departments (department, region, district_region, district_city, city_type, city, street, street_type, hous, post_index, partner, status, register, edrpou, address, partner_name, id_terminal, koatu, tax_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
RETURNING department`

	department := ""

	//fmt.Println("InsertDepartmentsFromFile...")
	arr := FileToArr(DataInDir + "departments.csv")[1:]
	sizeVec := len(arr[0])
	for _, vec := range arr {
		if len(vec) == sizeVec {
			err = db.QueryRow(sqlStatement, vec[0], vec[1], vec[2], vec[3], vec[4], vec[5], vec[6], vec[7], vec[8], vec[9], vec[10], vec[11], vec[12], vec[13], vec[14], vec[15], vec[16], vec[17], vec[18]).Scan(&department)
			if err != nil {
				countErr += 1
				fmt.Println("err", vec[0])
				continue
			} else {
				count += 1
			}
			//fmt.Println(department)
		} else {
			countErr += 1
			fmt.Println("bed vec size", vec[1])
		}
	}
	if countErr == 0 {
		info += fmt.Sprintf("отделения: %d", count)
	} else {
		info += fmt.Sprintf(">> dep err: %d ok: %d", countErr, count)
	}
	return info
}

func InsertTerminalsFromFile() string {
	info := ""
	count := 0
	count_err := 0
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println(err)
		//return err
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO terminals (department, termial, model, serial_number, date_manufacture, soft, producer, rne_rro, sealing, fiscal_number, oro_serial, oro_number, ticket_serial, ticket_1sheet, ticket_number, sending, books_arhiv, tickets_arhiv, to_rro, owner_rro, register, finish)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
RETURNING termial`

	termial := ""

	//fmt.Println("InsertTerminalsFromFile...")
	arr := FileToArr(DataInDir + "terminals.csv")[1:]
	sizeVec := len(arr[0])
	for _, vec := range arr {
		if true || len(vec) == sizeVec {
			err = db.QueryRow(sqlStatement, vec[0], vec[1], vec[2], vec[3], vec[4], vec[5], vec[6], vec[7], vec[8], vec[9], vec[10], vec[11], vec[12], vec[13], vec[14], vec[15], vec[16], vec[17], vec[18], vec[19], "", "").Scan(&termial)

			if err != nil {
				count_err += 1
				//fmt.Println("err", vec[1])
				continue
			} else {
				count += 1
			}
			//fmt.Println(termial)
		} else {
			//countErr += 1
			//fmt.Println("bed vec size", vec[1])
		}
	}
	if count_err == 0 {
		info += fmt.Sprintf("терминалы: %d", count)
	} else {
		info += fmt.Sprintf(">> term err: %d ok: %d", count_err, count)
	}
	return info
}

func InsertOtborFromFile() string {
	info := ""
	count := 0
	count_err := 0
	ok := true
	arr := FileToArr(DataInDir + "otbor.csv")[1:]

	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	for _, vec := range arr {
		sqlStatement := `
INSERT INTO otbor (term, dep)
VALUES ($1, $2)
RETURNING term`

		term := ""
		err = db.QueryRow(sqlStatement, vec[0], vec[1]).Scan(&term)
		count += 1
		if err != nil {
			ok = false
			fmt.Println(">> err insert otbor")
			count_err += 1
		}
	}
	if ok {
		info += fmt.Sprintf("success otbor %d", count)
		//fmt.Println("success")
	} else {
		info += fmt.Sprintf(">> otbor err: %d, ok: %d", count_err, count)
		//fmt.Println("errors")
	}
	return info
}

func ClearTableDepartments() {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		//return err
		panic(err)
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM departments;")
	if err != nil {
		panic(err)
		//return err
	}
	//fmt.Println("clear table departments")
	//return nil
}

func ClearTableTerminals() {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		//return err
		panic(err)
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM terminals;")
	if err != nil {
		panic(err)
		//return err
	}
	//fmt.Println("clear table terminals")
	//return nil
}

func InsertOtborOneRecord(vec []string) error {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return err
		//panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO otbor (term, dep)
VALUES ($1, $2)
RETURNING term`

	term := ""
	err = db.QueryRow(sqlStatement, vec[0], vec[1]).Scan(&term)
	if err != nil {
		//panic(err)
		return err
	}
	//fmt.Println("New record term is:", term)
	return nil
}
