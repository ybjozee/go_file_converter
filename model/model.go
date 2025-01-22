package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Reader interface {
	Read(filePath string) error
	GetContent() FileContent
}

type Writer interface {
	Write(content FileContent, fileName string, filePath string) error
}

type Converter interface {
	GetReader() Reader
	GetWriter() Writer
}

type FileContent struct {
	Headers []string
	Data    [][]string
}

type Conversion struct {
	gorm.Model
	UploadedFileName    string
	UploadedFileSize    string
	UploadedFileFormat  string
	ConvertedFileName   string
	ConvertedFileSize   string
	ConvertedFileFormat string
	ConversionDate      datatypes.Date
	NameToSaveAs        string
}

type TemplateContext struct {
	Error string
	Data  any
}
