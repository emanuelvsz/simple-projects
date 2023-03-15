package pdfGenerator

type PDFGeneratorInterface interface {
	Create(htmlFile string) (string, error)
}
