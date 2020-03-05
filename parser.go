package jrouter

import (
	"regexp"
	"strings"
)

const (
	//finderMatchPattern regexp to support paramters /something/{id}
	finderMatchPattern string = "(\\{[a-zA-Z\\d+:\\(\\)\\,]+\\})"
	//defaultParamMatcher regexp per parameter
	defaultParamMatcher string = "([a-zA-Z\\d+]+)"
)

//URLParser is the struct responsible to manage the url and how to parse it
type URLParser struct {
	Base           string
	Pattern        string
	PatternMatcher *regexp.Regexp
	Params         []Parameter
}

//Parameter is a struct for each paramater to be read
type Parameter struct {
	Matcher string
	Param   string
}

//NewURLParser Instance new struct
func NewURLParser() *URLParser {
	return &URLParser{}
}

//Analyze methos allow us analyze the url
func (p *URLParser) Analyze(url string) *URLParser {
	re := regexp.MustCompile(finderMatchPattern)

	params := re.FindAllString(url, -1)

	if params == nil {
		p.PatternMatcher = regexp.MustCompile(url)
		p.Base = url
	} else {

		var matchers []string

		for _, param := range params {
			paramStr := param[1 : len(param)-1]
			pa := Parameter{Param: paramStr, Matcher: defaultParamMatcher}
			p.Params = append(p.Params, pa)
			matchers = append(matchers, pa.Matcher)
		}

		urlParts := strings.Split(url, "{")

		p.Base = urlParts[0]
		urlPattern := p.Base + strings.Join(matchers, "/")
		p.PatternMatcher = regexp.MustCompile(urlPattern)

	}
	return p
}
