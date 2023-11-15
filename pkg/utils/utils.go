package utils

import (
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func ExtractTextFromPage(page *model.PdfPage) string {
	extractor, err := extractor.New(page)
	if err != nil {
		return ""
	}

	text, err := extractor.ExtractText()
	if err != nil {
		return ""
	}

	return text
}
