package converter

import (
	"app/model"
	"app/reader"
	"app/writer"
)

type CSVConverter struct{}

func (converter CSVConverter) GetReader() model.Reader {
	return &reader.CSVReader{}
}

func (converter CSVConverter) GetWriter() model.Writer {
	return &writer.CSVWriter{}
}
