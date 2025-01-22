package converter

import (
	"app/model"
	"errors"
)

func GetConverter(extension string) (model.Converter, error) {
	switch extension {
	case "csv":
		return &CSVConverter{}, nil
	case "json":
		return &JSONConverter{}, nil
	case "sql":
		return &SQLConverter{}, nil
	case "xlsx":
		return &XLSXConverter{}, nil
	default:
		return nil, errors.New("file format not currently supported")
	}
}

func GetSupportedFormats() []string {
	return []string{
		"csv",
		"json",
		"sql",
		"xlsx",
	}
}
