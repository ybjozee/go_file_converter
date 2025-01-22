package handler

import (
	"app/converter"
	"app/database"
	"app/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"gorm.io/datatypes"
)

func convertFile(w http.ResponseWriter, r *http.Request) {
	clearErrorMessage()

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	check(err)

	fileExtension := filepath.Ext(handler.Filename)
	outputFormat := r.FormValue("outputFormat")
	nameToSaveAs := r.FormValue("nameToSaveAs")
	inputFormat := strings.ReplaceAll(fileExtension, ".", "")

	if inputFormat == outputFormat {
		redirectWithErrorMessage("Input and output formats cannot be the same", w, r)
		return
	}

	tempFile, err := os.CreateTemp("temp", "upload-*"+fileExtension)
	check(err)

	fileName := tempFile.Name()
	defer func() {
		file.Close()
		tempFile.Close()
		os.Remove(fileName)
	}()

	_, err = io.Copy(tempFile, file)
	check(err)

	inputConverter, err := converter.GetConverter(inputFormat)
	if err != nil {
		redirectWithErrorMessage(err.Error(), w, r)
		return
	}

	reader := inputConverter.GetReader()
	err = reader.Read(fileName)
	check(err)

	outputConverter, err := converter.GetConverter(outputFormat)
	check(err)

	convertedFileName := fmt.Sprintf("%s_%d", nameToSaveAs, time.Now().Unix())
	content := reader.GetContent()
	err = outputConverter.GetWriter().Write(content, convertedFileName, "./output")
	check(err)

	output, _ := os.ReadFile(fmt.Sprintf("./output/%s.%s", convertedFileName, outputFormat))
	convertedFileSize := int64(len(output))

	database.Save(model.Conversion{
		UploadedFileName:    handler.Filename,
		UploadedFileSize:    readableFileSize(handler.Size),
		UploadedFileFormat:  inputFormat,
		ConvertedFileName:   convertedFileName,
		ConvertedFileSize:   readableFileSize(convertedFileSize),
		ConvertedFileFormat: outputFormat,
		ConversionDate:      datatypes.Date(time.Now()),
		NameToSaveAs:        nameToSaveAs,
	})

	w.Header().Set("Content-Disposition", "attachment; filename="+fmt.Sprintf("%s.%s", nameToSaveAs, outputFormat))
	w.Write(output)
}

func index(w http.ResponseWriter, r *http.Request) {
	context.Data = converter.GetSupportedFormats()
	renderTemplate(w, "template/index.html")
}

func getConversions(w http.ResponseWriter, r *http.Request) {
	clearErrorMessage()

	conversions, err := database.GetConversions()
	check(err)

	context.Data = conversions

	renderTemplate(w, "template/conversions.html")
}

func getConversion(w http.ResponseWriter, r *http.Request) {
	clearErrorMessage()

	id := r.PathValue("id")

	conversion, err := database.GetConversion(id)
	if err != nil {
		context.Error = err.Error()
		http.Redirect(w, r, "/conversions", http.StatusTemporaryRedirect)
		return
	}

	output, _ := os.ReadFile(fmt.Sprintf("./output/%s.%s", conversion.ConvertedFileName, conversion.ConvertedFileFormat))

	w.Header().Set("Content-Disposition", "attachment; filename="+fmt.Sprintf("%s.%s", conversion.NameToSaveAs, conversion.ConvertedFileFormat))
	w.Write(output)
}

func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/convert", convertFile)
	mux.HandleFunc("/conversions", getConversions)
	mux.HandleFunc("/conversion/result/{id}", getConversion)

	return mux
}
