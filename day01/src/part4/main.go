package main
// Save Иванов Иван Иванович\оnортопед\n 2024-04-13\n
import (
	"bufio"
	"fmt"
	"os"
)

type infoVisits struct {
	docSpecial string
	dateVisit  string
}

type UserNotFoundError struct {
	Name string
}

func (*UserNotFoundError) Error() string {
	return "user not found"
}

func main() {
	userDataBase := map[infoVisits]string{}
	var commands string
	exitProgram := false
	for exitProgram != true {
		fmt.Println("Нашите команду: Save, GetHistory, GetLastVisit")
		fmt.Scan(&commands)
		switch commands {
		case "Save":
			saveVisits(userDataBase)
		case "GetHistory":
			fmt.Println("Напишите ФИО в таком формате: Иванов Иван Иванович")
			nameUserGetHistory := readString()
			result, err := getHistory(userDataBase, nameUserGetHistory)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				printGetHistory(result)
			}
		case "GetLastVisit":
			fmt.Println("Напишите ФИО в таком формате: Иванов Иван Иванович")
			nameUserGetLast := readString()
			fmt.Println("Напишите специализацию врача")
			docSpecialGetLast := readString()
			result, err := getLastVisit(nameUserGetLast, docSpecialGetLast, userDataBase)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(result)
			}
		case "Exit":
			exitProgram = true
		}
	}
}

func readString() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		return line
	}
	return ""
}

func saveVisits(dataBase map[infoVisits]string) {
	fmt.Println("Напишите ФИО в таком формате: Иванов Иван Иванович")
	nameInput := readString()
	fmt.Println("Напишите специализацию врача")
	docInput := readString()
	fmt.Println("Напишите дату посещения в фомате: ГГГГ-ММ-ДД")
	dateInput := readString()
	dataBase[infoVisits{docSpecial: docInput, dateVisit: dateInput}] = nameInput
}

func getHistory(dataBase map[infoVisits]string, nameUser string) ([]string, error) {
	resultSlice := []string{}
	for k, v := range dataBase {
		if nameUser != v {
			return resultSlice, &UserNotFoundError{Name: nameUser}
		}
		resultSlice = append(resultSlice, k.docSpecial, k.dateVisit)
	}
	return resultSlice, nil
}

func printGetHistory(s1 []string) {
	k, j := 0, 1
	for i := range s1 {
		if i < len(s1)/2 {
			fmt.Printf("%s %s\n", s1[k], s1[j])
			k += 2
			j += 2
		}
	}
}

func getLastVisit(nameUser, special string, dataBase map[infoVisits]string) (string, error) {
	sliceDataVisit := []string{}
	for k, v := range dataBase{
		if v != nameUser{
			return "", &UserNotFoundError{Name: nameUser}
		}
		if k.docSpecial == special{
		sliceDataVisit = append(sliceDataVisit, k.dateVisit)
}else{
	continue
}
	}
	return sliceDataVisit[len(sliceDataVisit)-1], nil
}