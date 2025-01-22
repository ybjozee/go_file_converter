package reader

import (
	"app/model"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"os"
	"slices"
)

type JSONReader struct {
	content model.FileContent
}

func (reader *JSONReader) Read(filePath string) error {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var jsonObject []map[string]any
	if err := json.Unmarshal(f, &jsonObject); err != nil {
		return err
	}

	if len(jsonObject) == 0 {
		return errors.New("invalid json provided")
	}

	reader.content.Headers = slices.Collect(maps.Keys(jsonObject[0]))
	reader.content.Data = readJson(jsonObject, reader.content.Headers)

	return nil
}

func (reader *JSONReader) GetContent() model.FileContent {
	return reader.content
}

func readJson(jsonObject []map[string]any, headers []string) [][]string {
	output := [][]string{}
	for _, item := range jsonObject {
		data := []string{}
		for _, header := range headers {
			value := item[header]
			data = append(data, getJsonValue(value))
		}
		output = append(output, data)
	}
	return output
}

func getJsonValue(input any) string {
	switch value := input.(type) {
	case nil:
		return ""
	case string:
		return value
	default:
		return fmt.Sprintf("%v", value)
	}
}
