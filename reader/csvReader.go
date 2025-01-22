package reader

import (
	"app/model"
	"encoding/csv"
	"errors"
	"os"
)

type CSVReader struct {
	content model.FileContent
}

func (reader *CSVReader) Read(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return errors.New("Unable to read input file " + filePath)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return errors.New("Unable to parse file as CSV for " + filePath)
	}

	reader.content.Headers = records[0]
	reader.content.Data = records[1:]
	
	return nil
}

func (reader *CSVReader) GetContent() model.FileContent {
	return reader.content
}
