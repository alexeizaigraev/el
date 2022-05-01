// perevod
package people

import (
	"el/module"
	"fmt"
	"strings"
)

//var KassAll = MkKassShort()

func PerevodMain() module.Out6 {
	rez := ""
	inFname := module.DataInDir + "perevod.csv"
	outFname := module.DataOutDir + "OutPerevod.csv"

	OutText = ""
	Unfind = ""

	ColFioOne = 0
	ColDepOne = 1
	KassAll = MkKassShort()
	//KassAllHash = MkKassHash()
	text, _ := module.FileToVec(inFname)
	for _, line := range text {
		perevodProcess(line)
	}

	fmt.Println("\n\n" + OutText + "\n" + Unfind + "\n\n")
	strings.Trim(OutText, "\n")
	MyOut(Unfind+OutText, outFname)
	//CheckOutText(outText)
	if Unfind != "" {
		rez = "unfind"
		fmt.Println("\n\tunfind\n")
	} else {
		fmt.Println("\n\twell\n")
		rez = "well"
	}
	//module.OpenNote(outFname)
	return module.Out6{
		Kind:  "Перевод",
		Data:  strings.Split(Unfind+OutText, "/n"),
		Fname: outFname,
		Err:   "",
		Rez:   rez,
	}
}

func perevodProcess(line string) {
	LineSplit = strings.Split(line, ";")
	if len(LineSplit) < 3 {
		return
	}
	login := SearchLogin(MkOldDep())
	//login := SearchLoginHash()
	if LoginOk {
		for _, unit := range login {
			OutText += unit + "\t" + LineSplit[1] + " -> " + LineSplit[2] + "\n"
		}
	} else {
		Unfind += LineSplit[0] + "\t" + LineSplit[1] + " -> " + LineSplit[2] + "\n"
	}
}

/*
func perevodProcessHash(line string) {
	LineSplit = strings.Split(line, ";")
	if len(LineSplit) < 3 {
		return
	}
	//login := SearchLogin(MkOldDep())
	login := SearchLoginHash()
	if LoginOk {
		OutText += login + "\t" + LineSplit[1] + " -> " + LineSplit[2] + "\n"
	} else {
		Unfind += LineSplit[0] + "\t" + LineSplit[1] + " -> " + LineSplit[2] + "\n"
	}
}

*/
