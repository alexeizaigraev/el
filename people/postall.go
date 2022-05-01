// postall
package people

import (
	//"fmt"
	"el/module"
	"strings"
)

func PostAllMain() module.Out6 {
	outText := ""
	postOutPath := mkPostPath()

	postKassAll, _ := module.FileToVec(module.DataInDir + "kass_all.csv")

	outFName := postOutPath + "justin/OutPostAll.csv"
	TextToFile(mkKassAll(postKassAll, "justin"), outFName)
	outText += outFName + "\n"

	outFName = postOutPath + "justin/OutPostOtpuskaJust.csv"
	TextToFile(mkOtpuskAll(postKassAll, "justin"), outFName)
	outText += outFName + "\n"

	outFNameAllo := postOutPath + "allo/OutPostAllAllo.csv"
	TextToFile(mkKassAll(postKassAll, "allo"), outFNameAllo)
	outText += outFName + "\n"

	return module.Out6{
		Kind:  "Посталл",
		Data:  strings.Split(outText, "/n"),
		Fname: "",
		Err:   "",
		Rez:   "well",
	}
}

func mkKassAll(a []string, agent string) string {
	var out = "Логин;ФИО;Терминал\n"
	//a, _ := FileToVec(DataInDir + "kass_all.csv")
	for i := range a {
		splitLine := strings.Split(a[i], ";")
		if len(splitLine) > 3 && strings.Contains(splitLine[1], "true") && strings.Contains(splitLine[4], agent) {
			line := splitLine[0] + ";" + splitLine[2] + ";" + splitLine[3]
			out += line + "\n"
			//fmt.Println(splitLine[0] + ";" + splitLine[10])
		}
	}
	return out
}

func mkAllLogins(a []string, agent string) []string {
	var out []string
	//a, _ := FileToVec(DataInDir + "kass_all.csv")
	for i := range a {
		splitLine := strings.Split(a[i], ";")
		if len(splitLine) > 3 && strings.Contains(splitLine[1], "true") && strings.Contains(splitLine[4], agent) {
			out = append(out, splitLine[0])
		}
	}
	return out
}

func mkOtpuskAll(postKassAll []string, agent string) string {
	var out = "Логин;Начало отпуска;Конец отпуска;Уволен\n"
	myLogins := mkAllLogins(postKassAll, agent)
	myOtpuska, _ := module.FileToVec(module.DataInDir + "all_otpuska.csv")

	for i := range myOtpuska {
		splitOtpLine := strings.Split(myOtpuska[i], ";")
		//fmt.Println(splitOtpLine[0])
		for j := range myLogins {
			//fmt.Println(myLogins[j])
			if strings.Contains(myLogins[j], splitOtpLine[0]) {
				line := splitOtpLine[0] + ";" + splitOtpLine[1] + ";" + splitOtpLine[2] + ";" + splitOtpLine[3]
				out += line + "\n"
				//fmt.Println(line)
				break
			}
		}

	}
	return out
}

func mkPostPath() string {
	s, _ := module.FileToVec(module.DataConfigPath + "ConfigPostPath.txt")
	return s[0]
}
