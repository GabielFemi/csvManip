package gabby_csv

import (
	"encoding/csv"
	"os"
)

func ReadCsv(csvFile string) []Employee {
	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Instantiate reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comma = ';'

	csvData, err := reader.ReadAll()
	var emp Employee
	var employees []Employee
	for _, item := range csvData {


		emp.Name = item[0]
		emp.Email = item[1]
		emp.Country = item[2]

		employees = append(employees, emp)

	}
	return employees
}
