package htmlParser

type HTMLParserInterface interface {
	Create(templateName string, data interface{}) (string, error)
}
