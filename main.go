package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Identification struct {
	Letter string
	Number int
}

func main() {

	csvFile, err := os.Open("test.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	records, _ := csvReader.ReadAll()

	var jsonText = []byte(`[]`)

	var idents []Identification

	if err := json.Unmarshal([]byte(jsonText), &idents); err != nil {
		log.Println(err)
	}

	for _, record := range records {

		num, err := strconv.Atoi(string(record[0][2]))
		if err != nil {
			fmt.Println(err)
			return
		}

		idents = append(idents,
			Identification{Letter: string(record[0][0]), Number: num},
		)

	}

	result, err := json.Marshal(idents)
	if err != nil {
		log.Println(err)
	}

	os.WriteFile("test.json", result, 0644)

}
