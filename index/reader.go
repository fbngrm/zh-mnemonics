package index

import (
	"encoding/csv"
	"fmt"
	"os"
)

func New(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	index := make(map[string]string)

	for _, record := range records {
		if len(record) >= 3 {
			key := record[0]
			value := record[2]
			index[key] = value
		} else {
			fmt.Printf("skipping incomplete record: %v\n", record)
		}
	}
	return index, nil
}
