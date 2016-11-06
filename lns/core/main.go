package core

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type LunchAndSnack struct {
	Date  string
	Lunch string
	Snack string
}

// type LunchAndSnack struct {
// 	Date  time.Time
// 	Lunch []string
// 	Snack []string
// }

func ReadCsv() (tsvData [][]string) {
	// read the tsv file
	tsvFile, err := os.Open(os.Getenv("CSV_DATA_INPUT"))
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer tsvFile.Close()

	reader := csv.NewReader(tsvFile)
	// reader.Comma = '\t'         // our data src is in tsv :D
	reader.FieldsPerRecord = -1 //assuming records might have variable number of fields, so no check made

	tsvData, err = reader.ReadAll()
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
	return
}

func AppendToStructArray(tsvData *[][]string) (foodAll []LunchAndSnack) {
	var lns LunchAndSnack

	//S.N.	Date	Day	Lunch	Snack
	for _, each := range *tsvData {
		lns.Date = each[1]
		lns.Lunch = each[3]
		lns.Snack = each[4]

		foodAll = append(foodAll, lns)
	}
	return
}

func Export2JsonFile(foodAll *[]LunchAndSnack) {
	//convert to json
	jsondata, err := json.Marshal(*foodAll)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	//lets write to a file
	jsonFile, err := os.Create(os.Getenv("JSON_DATA_OUTPUT"))
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
}

func Csv2Json() {
	tsvData := ReadCsv()
	foodAll := AppendToStructArray(&tsvData)
	Export2JsonFile(&foodAll)
}

// func (f *LunchAndSnack) GetFoodItem() {
// 	foodAll := AppendToStructArray(&ReadCsv())

// }
