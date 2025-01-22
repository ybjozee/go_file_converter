package converter

import (
	"app/model"
	"app/reader"
	"app/writer"
)

type SQLConverter struct {
}

func (converter SQLConverter) GetReader() model.Reader {
	return &reader.SQLReader{}
}

func (converter SQLConverter) GetWriter() model.Writer {
	return &writer.SQLWriter{}
}
