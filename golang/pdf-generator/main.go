package main

import (
	"fmt"
	htmlParser "modulo/generators/htmlGenerator"
	"modulo/generators/pdfGenerator"
	"modulo/models"
)

func main() {
	h := htmlParser.New("tmp/html")
	wk := pdfGenerator.NewWKHTMLToPDF("tmp/pdf")

	dataHTML := models.Data{
		Name:        "Emanuel Vilela de Souza",
		Age:         19,
		Description: "Curto muito música e amo programar",
		Formation:   "Formado em Técnico em Informática no Instituto Federal de Alagoas(IFAL) e cursando Sistemas de Informação(IFAL)",
		CPF:         "105.539.714-09",
	}

	htmlGenerated, err := h.Create("./utils/templates/pdf.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("HTML gerado: ", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("PDF gerado: ", filePDFName)
}
