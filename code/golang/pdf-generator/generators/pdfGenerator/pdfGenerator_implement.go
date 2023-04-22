package pdfGenerator

import (
	"fmt"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wkhtml struct {
	rootPath string
}

func NewWKHTMLToPDF(rootPath string) PDFGeneratorInterface {
	return &wkhtml{rootPath: rootPath}
}

func (w *wkhtml) Create(htmlFile string) (string, error) {
	file, err := os.Open(htmlFile)
	if err != nil {
		return "erro aqui 1", nil
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(file))
	if err != nil {
		fmt.Println("Deu tudo errado: ", err)
		return "", err
	}

	if err := pdfg.Create(); err != nil {
		fmt.Println("Erro aqui 2")
		return "erro aqui 2", err
	}

	fmt.Println("Ews")

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	if err := pdfg.WriteFile(fileName); err != nil {
		return "erro aqui 3", err
	}

	return fileName, nil
}
