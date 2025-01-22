package converter

import (
	"app/model"
	"app/reader"
	"app/writer"
)

type JSONConverter struct{}

func (converter JSONConverter) GetReader() model.Reader {
	return &reader.JSONReader{}
}

func (converter JSONConverter) GetWriter() model.Writer {
	return &writer.JSONWriter{}
}
