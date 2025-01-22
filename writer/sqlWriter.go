package writer

import (
	"app/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SQLWriter struct {
}

func (writer *SQLWriter) Write(content model.FileContent, fileName string, filePath string) error {
	sqlContent := fmt.Sprintf("%s%s", getInsertStatement(fileName, content.Headers), writeValues(content.Data))

	file, err := os.Create(fmt.Sprintf("%s/%s.sql", filePath, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(sqlContent)
	if err != nil {
		return err
	}

	return nil
}

func getInsertStatement(fileName string, keys []string) string {
	keyString := strings.Join(keys, "`,`")
	return fmt.Sprintf("INSERT INTO `%s` (`%s`) VALUES \n", fileName, keyString)
}

func writeValues(values [][]string) string {
	sqlString := ""

	for _, row := range values {
		columnString := ""
		for _, column := range row {
			columnString += getFormattedColumn(column) + ", "
		}
		sqlString += fmt.Sprintf("\n(%s),", columnString)
	}

	sqlString += ";"
	sqlString = strings.ReplaceAll(sqlString, ", )", ")")

	return strings.Replace(sqlString, ",;", ";", 1)
}

func getFormattedColumn(value string) string {
	if value == "" {
		return "NULL"
	}

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return value
	}

	if _, err := strconv.Atoi(value); err == nil {
		return value
	}

	return fmt.Sprintf("\"%s\"", value)
}
