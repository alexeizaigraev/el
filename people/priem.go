// priem
package people

import (
	"el/module"
	"fmt"
	"strings"
)

func PriemMain() module.Out6 {
	//OutText = ""
	ColFioOne = 0
	ColDepOne = 1
	inFname := module.DataInDir + "priem.csv"
	outFname := module.DataOutDir + "OutPriem.csv"
	OutText = "Логин;Пароль;ФИО;Почта;Телефон;Агент;Терминал\n"
	text, _ := module.FileToVec(inFname)

	for _, line := range text {
		priemProcess(line)
	}

	MyOut(OutText, outFname)

	return module.Out6{
		Kind:  "Приём",
		Data:  strings.Split(OutText, "/n"),
		Fname: outFname,
		Err:   "",
		Rez:   "well",
	}
}

func printer(c chan string) {
	s := <-c
	OutText += s
	fmt.Println(s)
}

func priemProcess(line string) {
	LineSplit = strings.Split(line, ";")
	//fmt.Println(LineSplit[0])
	name2 := MkSurname() +
		string([]rune(MkFirstName())[:2]) +
		string([]rune(MkLastName())[:2])
	name := strings.ReplaceAll(name2, "-", "")
	login := login(name)
	fio := MkShortName()

	dep := MkOldDep()
	//login += dep[len(dep)-3:]
	term := dep + "1"

	agent_sign := dep[:3]
	agent_dat := prichindal(agent_sign)
	split_ag_dat := strings.Split(agent_dat, ";")
	agent := split_ag_dat[1]
	mail := split_ag_dat[2]
	phone := "380999999999"
	if strings.Contains(LineSplit[3], "@") {
		mail = LineSplit[3]
	}
	if strings.Contains(LineSplit[2], "0") && strings.Contains(LineSplit[2], "-") && len(strings.Replace(LineSplit[2], "-", "", -1)) == 10 {
		phone = "38" + strings.Replace(LineSplit[2], "-", "", -1)
	}
	pasp_seria := "id"
	if len(LineSplit[4]) > 1 {
		pasp_seria = LineSplit[4]
	}
	pasp_number := "999999"
	if len(LineSplit[5]) > 4 {
		pasp_number = LineSplit[5]
	}
	pasp_vydan := "99999"
	if len(LineSplit[7]) > 4 {
		pasp_vydan = LineSplit[7]
	}
	date := "2020-01-01"
	if strings.Contains(LineSplit[6], ".") && len(strings.Replace(LineSplit[6], ".", "", -1)) == 8 {
		split_date := strings.Split(LineSplit[6], ".")
		date = string(split_date[2][:4] + "-" +
			split_date[1] + "-" + split_date[0])
	}
	login += pasp_number[len(pasp_number)-4:]
	outLine := login + ";" + login + ";" + fio + ";" +
		mail + ";" + phone + ";" + agent + ";" + term + ";" +
		pasp_seria + ";" + pasp_number + ";" + pasp_vydan + ";" + date + "\n"

	OutText += outLine
}

func login(word string) string {
	my_login := ""
	low_word := strings.ToLower(word)
	rune_word := []rune(low_word)
	//ua := [...]string{"а", "б", "в", "г", "ґ", "д", "е", "є", "ж", "з", "и", "і", "ї", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ь", "ю", "я"}
	//lat := [...]string{"a", "b", "v", "g", "gh", "d", "e", "je", "zh", "z", "i", "i", "ji", "j", "k", "l", "m", "n", "o", "p", "r", "s", "t", "u", "f", "h", "z", "sh", "sch", "s", "j", "yu", "ya"}
	d := map[string]string{
		"а": "a",
		"б": "b",
		"в": "v",
		"г": "g",
		"ґ": "gh",
		"д": "d",
		"е": "e",
		"є": "je",
		"ж": "zh",
		"з": "z",
		"и": "i",
		"і": "i",
		"ї": "ji",
		"й": "j",
		"к": "k",
		"л": "l",
		"м": "m",
		"н": "n",
		"о": "o",
		"п": "p",
		"р": "r",
		"с": "s",
		"т": "t",
		"у": "u",
		"ф": "f",
		"х": "h",
		"ц": "z",
		"ч": "sh",
		"ш": "sch",
		"щ": "s",
		"ь": "j",
		"ю": "yu",
		"я": "ya",
	}
	for _, cha := range rune_word {
		my_login += d[string(cha)]
	}
	return my_login
}

func prichindal(znak string) string {

	a, _ := module.FileToVec(module.DataConfigPath + "comon_data.csv")
	agent_data := "Error agent_data"
	for i := range a {
		split_line := strings.Split(a[i], ";")

		if strings.Contains(split_line[0], znak) {
			agent_data = a[i]
			break
		}
	}
	return strings.Trim(agent_data, "\n")
}
