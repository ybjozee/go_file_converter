package writer

import (
	"app/model"
	"fmt"
	"github.com/xuri/excelize/v2"
)

type XLSXWriter struct {
}

func (writer *XLSXWriter) Write(content model.FileContent, fileName string, filePath string) error {
	file := excelize.NewFile()

	defer func() error {
		if err := file.Close(); err != nil {
			return err
		}
		return nil
	}()

	if err := file.SetSheetName("Sheet1", fileName); err != nil {
		return err
	}

	index, err := file.GetSheetIndex(fileName)
	if err != nil {
		return err
	}

	file.SetActiveSheet(index)

	borderStyle, err := file.NewStyle(getBorderStyle())
	if err != nil {
		return err
	}

	rowIndex := 1
	writeHeader(file, content.Headers, rowIndex, fileName, borderStyle)
	rowIndex++

	writeData(file, content.Data, rowIndex, fileName, borderStyle)

	if err := file.SaveAs(fmt.Sprintf("%s/%s.xlsx", filePath, fileName)); err != nil {
		return err
	}

	return nil
}

func getBorderStyle() *excelize.Style {
	return &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	}
}

func writeHeader(file *excelize.File, headers []string, rowIndex int, fileName string, borderStyle int) {
	for index, value := range headers {
		column := getColumnFromNumber(index)
		cell := fmt.Sprintf("%s%d", column, rowIndex)
		file.SetCellValue(fileName, cell, value)
		file.SetCellStyle(fileName, cell, cell, borderStyle)
		file.SetColWidth(fileName, column, column, 40)
	}
}

func writeData(file *excelize.File, data [][]string, rowIndex int, fileName string, borderStyle int) {
	for _, value := range data {
		for cellIndex, cellData := range value {
			cell := fmt.Sprintf("%s%d", getColumnFromNumber(cellIndex), rowIndex)
			file.SetCellValue(fileName, cell, cellData)
			file.SetCellStyle(fileName, cell, cell, borderStyle)
		}
		rowIndex++
	}
}

func getColumnFromNumber(number int) string {
	letter := rune(65 + (number % 26))
	remainder := number / 26
	if remainder > 0 {
		return fmt.Sprintf("%s%c", getColumnFromNumber(remainder-1), letter)
	}
	return fmt.Sprintf("%c", letter)
}
