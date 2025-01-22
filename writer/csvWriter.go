package writer

import (
	"app/model"
	"encoding/csv"
	"fmt"
	"os"
)

type CSVWriter struct {
}

func (writer *CSVWriter) Write(content model.FileContent, fileName string, filePath string) error {
	file, err := os.Create(fmt.Sprintf("%s/%s.csv", filePath, fileName))
	if err != nil {
		return err
	}

	defer file.Close()

	fileWriter := csv.NewWriter(file)

	if err := fileWriter.Write(content.Headers); err != nil {
		return err
	}

	for _, data := range content.Data {
		if err := fileWriter.Write(data); err != nil {
			return err
		}
	}

	fileWriter.Flush()

	if err := fileWriter.Error(); err != nil {
		return err
	}

	return nil
}
