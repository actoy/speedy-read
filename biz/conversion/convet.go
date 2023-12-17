package conversion

import (
	"speedy/read/biz/conversion/api"
	"speedy/read/biz/conversion/api/generated"
)

var (
	articleConvert api.ArticleConvert = &generated.ArticleConvertImpl{}
)

var (
	ArticleDOToThrift = articleConvert.ArticleDOToThrift
	ArticleThriftToDO = articleConvert.ArticleThriftToDO
)
