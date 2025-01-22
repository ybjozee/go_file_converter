package converter

import (
	"app/model"
	"app/reader"
	"app/writer"
)

type XLSXConverter struct {
}

func (converter XLSXConverter) GetReader() model.Reader {
	return &reader.XLSXReader{}
}

func (converter XLSXConverter) GetWriter() model.Writer {
	return &writer.XLSXWriter{}
}
