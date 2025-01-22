package writer

import (
	"app/model"
	"encoding/json"
	"fmt"
	"os"
)

type JSONWriter struct {
}

func (writer *JSONWriter) Write(content model.FileContent, fileName string, filePath string) error {
	var jsonContent []map[string]any = make([]map[string]any, 0)

	for _, data := range content.Data {
		entry := make(map[string]any, 0)
		for index, value := range data {
			key := content.Headers[index]
			entry[key] = value
		}
		jsonContent = append(jsonContent, entry)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.json", filePath, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(jsonContent); err != nil {
		return err
	}

	return nil
}
