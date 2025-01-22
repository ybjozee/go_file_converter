package reader

import (
	"app/model"
	"errors"
	"os"
	"regexp"
	"strings"
)

type SQLReader struct {
	content model.FileContent
}

func (reader *SQLReader) Read(filePath string) error {

	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	sql := string(f)
	content := strings.Split(sql, "VALUES")
	if len(content) > 2 {
		return errors.New("invalid SQL provided")
	}

	reader.content.Headers = readHeader(content[0])
	reader.content.Data = readSQL(content[1])

	return nil
}

func (reader *SQLReader) GetContent() model.FileContent {
	return reader.content
}

func readHeader(header string) []string {
	insertRegex := regexp.MustCompile("INSERT INTO [A-Za-z`]*")
	withoutInsert := insertRegex.ReplaceAllString(header, "")

	otherCharactersRegex := regexp.MustCompile("[;`()' ]")
	withoutOtherCharacters := otherCharactersRegex.ReplaceAllString(withoutInsert, "")

	return strings.Split(withoutOtherCharacters, ",")
}

func readSQL(content string) [][]string {
	output := [][]string{}
	data := strings.Split(formatContent(content), "),")

	for _, datum := range data {
		output = append(output, strings.Split(datum, ","))
	}

	return output
}

func formatContent(content string) string {
	withoutSpecialCharacters := regexp.MustCompile("['('\n]")
	formattedContent := withoutSpecialCharacters.ReplaceAllString(content, "")

	withoutNull := regexp.MustCompile("NULL")
	formattedContent = withoutNull.ReplaceAllString(formattedContent, "")
	formattedContent = strings.ReplaceAll(formattedContent, ", ", ",")

	return strings.Replace(formattedContent, ");", "", 1)
}
