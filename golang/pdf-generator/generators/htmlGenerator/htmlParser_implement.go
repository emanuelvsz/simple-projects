package htmlParser

import (
	"fmt"
	"html/template"
	msgs "modulo/utils"
	"os"

	"github.com/google/uuid"
)

type htmlStruct struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
	return &htmlStruct{rootPath: rootPath}
}

func (a *htmlStruct) Create(templateName string, data interface{}) (string, error) {
	templateGenerator, err := template.ParseFiles(templateName)
	if err != nil {
		return msgs.GenerateTemplateError, nil
	}
	fileName := a.rootPath + "/" + uuid.New().String() + ".html" // da um nome unico ao arquivo 

	fmt.Println(fileName)

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return msgs.CreateArchiveError, nil
	}

	err = templateGenerator.Execute(fileWriter, data)
	if err != nil {
		return msgs.ExecuteArchiveError, nil
	}

	return fileName, nil
}
