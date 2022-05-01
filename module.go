// electrum project main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	//"path/filepath"

	"strings"
)

type Out6 struct {
	User  string
	Kind  string
	Data  []string
	Fname string
	Err   string
	Rez   string
}

var (
	ConnStr = "user=postgres password=postgres dbname=drm_go sslmode=disable"

	DataPath          = mkDataPath()
	DataInDir         = DataPath + "InData/"
	DataOutDir        = DataPath + "OutData/"
	DataConfigPath    = DataPath + "Config/"
	DataConfigDirPath = DataPath + "ConfigDir/"
	KabinetPath       = mkKabinetPath()

	OutText = ""
)

func GetUserName(r *http.Request) string {
	session, err := r.Cookie("session_id")
	loggedIn := (err != http.ErrNoCookie)
	if loggedIn {
		fmt.Println(session.Value)
		return session.Value
	}
	return "anonimus"
}

func mkDataPath() string {
	s, err := FileToVec("Config/ConfigDataPath.txt")
	if err != nil {
		Sos("mkDataPath", "Config/ConfigDataPath.txt")
	}
	return s[0]
}

func mkGdrivePath() string {
	s, err := FileToVec(DataConfigPath + "ConfigGdrivePath.txt")
	if err != nil {
		Sos("mkGdrivePath", DataConfigPath+"ConfigGdrivePath.txt")
	}
	return s[0]
}

func mkKabinetPath() string {
	s, err := FileToVec(DataConfigPath + "ConfigKabinetPath.txt")
	if err != nil {
		Sos("mkGdrivePath", DataConfigPath+"ConfigKabinetPath.txt")
	}
	return s[0]
}

func mkComonDataPath() string { return DataConfigPath + "comon_data.csv" }

func Sos(kind string, text string) {
	var name string
	fmt.Println("\nerr", kind, text)
	fmt.Print(" Exit [Enter] -> ")
	fmt.Fscan(os.Stdin, &name)
	//MenuMain()
	//Menu()
	panic(fmt.Sprintf(" panic\nGood By"))
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func FileToVec(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		//return nil, err
		Sos("FileToVec", path)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line2 := strings.Trim(scanner.Text(), " ")
		line3 := strings.ReplaceAll(line2, "\n", "")
		if len(line3) > 3 {
			//fmt.Println(line3)
			lines = append(lines, line3)
		}

	}
	return lines, scanner.Err()
}

func FileToText(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		//return nil, err
		Sos("FileToText", path)
	}
	defer file.Close()
	lines := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += scanner.Text()
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func VecToFile(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		//return err
		Sos("VecToFile", path)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func TextToFile(out_text string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		//return err
		Sos("TextToFile", path)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	fmt.Fprintln(w, out_text)
	fmt.Println(path)
	return w.Flush()
}

func ArrToText(arr [][]string) string {
	outText := ""
	line := ""
	for _, v := range arr {
		line = strings.Join(v, ";")
		outText += line + "\n"
	}
	return outText
}

func ArrToFile(arr [][]string, fName string) {
	text := ArrToText(arr)
	TextToFile(text, fName)
}

func FileToArr(fname string) [][]string {
	var out [][]string
	vectors, _ := FileToVec(fname)
	for _, vector := range vectors {
		sl := strings.Split(vector, ";")
		out = append(out, sl)
	}
	return out
}

func VecToHash(head []string, vec []string) map[string]string {
	out := make(map[string]string)
	for i := 0; i < len(head) && i < len(vec); i++ {
		out[head[i]] = vec[i]
	}
	return out
}

func HashToStr(hash map[string]string) string {
	line := ""
	for key := range hash {
		line += hash[key] + ";"
	}
	return line
}

func ArrToHashTab(head []string, arr [][]string, key int) map[string]map[string]string {
	out := make(map[string]map[string]string)
	lineMap := make(map[string]string)
	for _, line := range arr {
		lineMap = VecToHash(head, line)
		out[lineMap[head[key]]] = lineMap
	}
	return out
}

func FileToHashTab(fname string, key int) map[string]map[string]string {
	a := FileToArr(fname)
	head := a[0]
	data := a[1:]
	out := ArrToHashTab(head, data, key)
	return out
}

func HashToFile(head []string, hash map[string]map[string]string, fname string) {
	text := strings.Join(head, ";") + "\n"

	for keyBig := range hash {
		hh := hash[keyBig]
		line := ""
		for _, key := range head {
			line += hh[key] + ";"
		}
		text += line + "\n"
	}

	TextToFile(text, fname)
}

func MKHead(fname string) []string { return FileToArr(fname)[0] }

func mkComonData(colNum int) map[string]string {
	h := make(map[string]string)

	a := FileToArr(mkComonDataPath())
	for _, vec := range a {
		h[vec[0]] = vec[colNum]
	}
	return h
}

func MkComonData(colNum int) map[string]string {
	h := make(map[string]string)

	a := FileToArr(mkComonDataPath())
	for _, vec := range a {
		h[vec[0]] = vec[colNum]
	}
	return h
}

// Copy a file
func CopyFile(inFname string, ofName string) {
	//Read all the contents of the  original file
	bytesRead, err := ioutil.ReadFile(inFname)
	if err != nil {
		log.Fatal(err)
	}

	//Copy all the contents to the desitination file
	err = ioutil.WriteFile(ofName, bytesRead, 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" copy ", ofName)
}

func ColKey(hh map[string]map[string]string, myKey string) string {
	s := make(map[string]string)
	for key := range hh {
		line := hh[key]
		s[line[myKey]] = ""
	}

	var listKey []string
	for kKey := range s {
		listKey = append(listKey, kKey)
	}

	for i := range listKey {
		if listKey[i] == "" {
			continue
		}
		fmt.Println("\t", i, listKey[i])
	}

	fmt.Print("\n -> ")
	var input int
	fmt.Fscan(os.Stdin, &input)

	return listKey[input]

}

func CheckOutText(text string) {
	if "" == text {
		fmt.Println("\n\n\tempty\n")
	} else {
		fmt.Println("\n\n\twell\n")
	}
}

func MoveFile(pathOld string, pathNew string) bool {
	err := os.Rename(pathOld, pathNew)

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println(" move ", pathNew)
		return true
	}
}

func MoveFileOtherDrive(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println(" Err MoveFileOtherDrive (Open)", sourcePath, " -> ", destPath)
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		fmt.Println(" Err MoveFileOtherDrive (Create)", sourcePath, " -> ", destPath)
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		fmt.Println(" Err MoveFileOtherDrive (Copy)", sourcePath, " -> ", destPath)
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		fmt.Println(" Err MoveFileOtherDrive (Remove)", sourcePath, " -> ", destPath)
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	fmt.Println(" move ", destPath)
	return nil
}

func GetFnamesOneFolder(dir string) []string {
	var fNames []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		Sos("GetFnamesOneFolder", dir)
	}

	for _, file := range files {
		if !file.IsDir() {
			fNames = append(fNames, dir+file.Name())
			//fmt.Println(" get ", dir+file.Name())
		} else {
			//fmt.Println(" file.IsDir ", file.Name())
		}
	}
	return fNames
}

func GetFnamesManyFolder(dirs []string) []string {
	var fNames []string
	for _, dir := range dirs {
		vec := GetFnamesOneFolder(dir)
		fNames = append(fNames, vec...)
	}
	return fNames
}

func PathExists(path string) bool {
	flag := true
	_, err := os.Stat(path)
	if err != nil {
		flag = false
	}
	return flag
}

func FolderOk(path string) {
	if PathExists(path) {
		return
	} else {
		os.Mkdir(path, os.ModePerm)
		fmt.Println("mk folger", path)
	}
}

func NowDateKabinet() string {
	today := time.Now()
	dateFull := strings.ReplaceAll(today.String(), "-", "")
	yy := string([]rune(dateFull)[:4])
	mm := string([]rune(dateFull)[4:6])
	dd := string([]rune(dateFull)[6:8])

	return dd + mm + yy
}

func MkNatasha() map[string]string {
	h := make(map[string]string)
	sign := "Відділення № "
	a := FileToArr(DataInDir + "natasha.csv")
	for _, line := range a {
		for _, unit := range line {
			if strings.Contains(unit, sign) {
				key := strings.Trim((strings.Split(unit, sign)[1]), " ")
				h[key] = ""
			}
		}
	}
	return h
}

func OpenNote(fname string) {

	path, err := exec.LookPath("notepad.exe")
	if err != nil {
		fmt.Println("Файл не найден")
	}
	fmt.Printf("Доступ к файлу %s\n", path)
	//cmd := exec.Command(path)
	cmd := exec.Command("notepad.exe", fname)
	cmd.Run()

}
