// app_01 project main.go
package people

import (
	"el/module"
	"strings"
)

var (
	KassAll = MkKassShort()
	//KassAllHash       = MkKassHash()
	//Line      string
	LineSplit []string

	ColFioOne int
	ColFioTwo int
	ColDepOne int
	ColDepTwo int

	LoginOk bool

	OutText = ""
	Unfind  = ""
)

func MkFio() string {
	l1 := strings.Trim(LineSplit[ColFioOne], " ")
	l2 := strings.ReplaceAll(l1, "  ", " ")
	sl := strings.Split(l2, " ")
	surName := sl[0]
	firstName := sl[1]
	lastName := sl[1]
	if len(sl) > 2 {
		lastName = sl[2]
	}
	fioOk := surName + " " + firstName + " " + lastName
	return fioOk
}

func MkSurname() string {
	return strings.Split(MkFio(), " ")[0]
}

func isUpper(cha string) bool {
	var flag bool
	if strings.ToUpper(cha) == cha {
		flag = true
	}
	return flag
}

func mkFioWhite(fff string) string {
	fffClear := strings.ReplaceAll(fff, "  ", " ")
	fffClear2 := strings.Trim(fffClear, " ")
	fs := strings.Split(fffClear2, " ")

	surn := fs[0]
	out := surn
	other0 := strings.Join(fs[1:], "")
	other1 := strings.ReplaceAll(other0, ".", "")
	other := strings.ReplaceAll(other1, " ", "")

	for _, leter := range []rune(other) {
		leterStr := string(leter)
		if isUpper(leterStr) {
			out += leterStr
		}
	}
	return out
}

func MkFirstName() string {
	return strings.Split(MkFio(), " ")[1]
}

func MkLastName() string {
	return strings.Split(MkFio(), " ")[2]
}

func MkInOneDot() string {
	return string([]rune(MkFirstName())[0]) + "."
}

func MkInTwoDot() string {
	return string([]rune(MkLastName())[0]) + "."
}

func MkShortName() string {
	return MkSurname() + " " + MkInOneDot() + " " + MkInTwoDot()
}

func MkOldDep() string {
	var dep string
	depline := LineSplit[ColDepOne]
	//fmt.Println(depline)
	if strings.Contains(depline, "№") {
		dep1 := strings.Split(depline, "№")[1]
		dep = strings.ReplaceAll(dep1, " ", "")
	} else {
		dep = strings.ReplaceAll(depline, " ", "")
	}
	return dep
}

func MkNewDep() string {
	return LineSplit[2]
}

func MyOut(outText string, outFname string) {
	module.TextToFile(outText, outFname)
	//fmt.Println("\nOk -> ", outFname, "\n")
}

func TextToFile(outText, outFname string) {
	panic("unimplemented")
}

func MkKassShort() [][3]string {
	var out [][3]string
	a, _ := module.FileToVec(module.DataInDir + "kass_all.csv")
	var outVec [3]string
	for _, v := range a {
		vec := strings.Split(v, ";")
		if len(vec) > 4 && strings.Contains(vec[1], "true") {
			fio := mkFioWhite(vec[2])
			outVec[0], outVec[1], outVec[2] = vec[0], fio, vec[3]
			//outVec := outVec0[:]
			out = append(out, outVec)
		}
	}
	return out
}

func SearchLoginDeep(parDep string, nama string) []string {
	parDep = string([]rune(MkOldDep())[:3])
	LoginOk = false
	logins := make([]string, 0)
	for _, kassLine := range KassAll {
		depDep := string([]rune(kassLine[2])[1:4])
		if strings.Contains(depDep, parDep) &&
			strings.Contains(kassLine[1], nama) {
			logins = append(logins, kassLine[0])
		}
	}

	if 1 == len(logins) {
		LoginOk = true
		return logins
	}
	out := make([]string, 0)
	return out
}

func SearchLogin(parDep string) []string {
	parDep = MkOldDep()
	LoginOk = false
	nama := mkFioWhite(MkFio())
	logins := make([]string, 0)
	for _, kassLine := range KassAll {
		if strings.Contains(kassLine[2], parDep) &&
			strings.Contains(kassLine[1], nama) {
			logins = append(logins, kassLine[0])
			LoginOk = true
			//fmt.Println("kassLine[2]", kassLine[2])

		}
	}

	if LoginOk {
		return logins
	}

	return SearchLoginDeep(parDep, nama)
}

/*
func MkDeps(deps string) []string {
	var myDeps []string
	s := strings.ReplaceAll(deps, "[", "")
	s = strings.ReplaceAll(s, "]", "")
	s = strings.ReplaceAll(s, " ", "")
	if strings.Contains(s, ",") {
		depSplit := strings.Split(s, ",")
		for _, unit := range depSplit {
			myDeps = append(myDeps, string([]rune(unit)[:7]))
		}
	} else {
		myDeps = append(myDeps, string([]rune(s)[:7]))
	}
	return myDeps
}

func MkKassHash() map[string]string {
	outHash := make(map[string]string)
	a, _ := FileToVec(DataInDir + "kass_all.csv")
	for _, v := range a {
		vec := strings.Split(v, ";")
		if len(vec) > 3 && strings.Contains(vec[1], "true") {
			fio := mkFioWhite(vec[2])
			for _, dep := range MkDeps(vec[3]) {
				key := fio + dep
				outHash[key] = vec[0]
				//fmt.Println(key, vec[0])
			}

		}
	}
	return outHash
}

func SearchLoginHashDeep() string {
	LoginOk = false
	count := 0
	login := ""
	parDep := string([]rune(MkOldDep())[:3])
	key := mkFioWhite(LineSplit[ColFioOne]) + parDep
	for keyReal := range KassAllHash {
		if strings.Contains(keyReal, key) {
			count += 1
			login = KassAllHash[keyReal]
		}
	}

	if 1 == count {
		LoginOk = true
		return login
	}

	return ""
}

func SearchLoginHash() string {
	LoginOk = false
	key := mkFioWhite(LineSplit[ColFioOne]) + MkOldDep()
	login, ok := KassAllHash[key]
	if !ok {
		return SearchLoginHashDeep()
	} else {
		LoginOk = true
		return login
	}
}

*/
