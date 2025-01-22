package reader

import (
	"app/model"
	"github.com/xuri/excelize/v2"
)

type XLSXReader struct {
	content model.FileContent
}

func (reader *XLSXReader) Read(filePath string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	firstSheet := f.WorkBook.Sheets.Sheet[0].Name
	content, err := f.GetRows(firstSheet)

	if err != nil {
		return err
	}

	reader.content.Headers = content[0]
	reader.content.Data = content[1:]

	return nil
}

func (reader *XLSXReader) GetContent() model.FileContent {
	return reader.content
}
