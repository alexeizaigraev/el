// otpusk
package people

import (
	"el/module"
	"fmt"
	"strings"
)

func OtpuskMain() module.Out6 {
	rez := ""
	ColFioOne = 0
	ColDepOne = 3
	inFnameOtpuskUvol := module.DataInDir + "otpusk_uvol.csv"
	outFname := module.DataOutDir + "OutOtpuskUvol.csv"

	OutText = ""
	Unfind := ""

	KassAll = MkKassShort()
	//KassAllHash = MkKassHash()
	text, _ := module.FileToVec(inFnameOtpuskUvol)
	for _, line := range text {
		otpuskProcess(line)
	}
	MyOut(Unfind+OutText, outFname)
	//CheckOutText(outText)
	if Unfind != "" {
		rez = "unfind"
		fmt.Println("\n\tunfind\n")
	} else {
		fmt.Println("\n\twell\n")
		rez = "well"
	}
	module.OpenNote(outFname)
	return module.Out6{
		Kind:  "Отпуск",
		Data:  strings.Split(Unfind+OutText, "/n"),
		Fname: outFname,
		Err:   "",
		Rez:   rez,
	}
}

/*
func otpuskProcessHash(line string) {
	LineSplit = strings.Split(line, ";")
	if len(LineSplit) < 3 {
		return
	}

	//login := SearchLogin(MkOldDep())
	login := SearchLoginHash()
	if LoginOk {
		OutText += login + ";" + dates() + "\n"
	} else {
		Unfind += LineSplit[0] + "\t" + LineSplit[3] + "\t" + ";" + dates() + "\n"
	}
}
*/

func otpuskProcess(line string) {
	LineSplit = strings.Split(line, ";")
	if len(LineSplit) < 3 {
		return
	}

	login := SearchLogin(MkOldDep())
	//login := SearchLoginHash()
	if LoginOk {
		for _, unit := range login {
			OutText += unit + ";" + dates() + "\n"
		}

	} else {
		Unfind += LineSplit[0] + "\t" + LineSplit[3] + "\t" + ";" + dates() + "\n"
	}
}

func dates() string {
	// Начало отпуска:
	date_1 := ""
	// Конец отпуска:
	date_2 := ""
	// Уволен:
	date_3 := ""
	// Активный с:
	date_4 := "2020-01-01 00:00:01"
	// Активный до:
	//date_5 := "2040-01-01 00:00:01"
	date_5 := ""

	//Otpusk
	if strings.Contains(LineSplit[2], "20") {
		b1 := strings.Split(LineSplit[1], ".")
		date_1 = b1[2] + "-" + b1[1] + "-" + b1[0] +
			" " + "00:00:01"
		b2 := strings.Split(LineSplit[2], ".")
		date_2 = b2[2] + "-" + b2[1] + "-" + b2[0] +
			" " + "23:59:59"
		date_3 = ""
		date_5 += "2050-01-01 00:00:01"
	} else {
		date_1 = ""
		date_2 = ""
		uvol := strings.Split(LineSplit[1], ".")
		date_3 = uvol[2] + "-" + uvol[1] + "-" + uvol[0] +
			" " + "23:59:59"
		date_5 += date_3
	}
	ddd := date_1 + ";" + date_2 + ";" +
		date_3 + ";" + date_4 + ";" + date_5
	return ddd
}
