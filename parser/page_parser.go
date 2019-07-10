package parser

type PageParser interface {
	PageProductSearchPage(pageURL string) ([]string, error)
}

type pageParserImpl struct {
}

func NewPageParser() PageParser {
	return &pageParserImpl{}
}

func (p *pageParserImpl) PageProductSearchPage(pageURL string) ([]string, error) {

}
