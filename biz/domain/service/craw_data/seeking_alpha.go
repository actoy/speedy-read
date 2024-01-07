package craw_data

type SeekingAlpha struct {
	Channel SeekingChannel `xml:"channel"`
}

type SeekingChannel struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description string        `xml:"description"`
	Item        []SeekingItem `xml:"item"`
}

type SeekingItem struct {
	Title      string         `xml:"title"`
	ArticleUrl string         `xml:"link"`
	PubDate    string         `xml:"pubDate"`
	AuthorName string         `xml:"author_name"`
	Stock      []SeekingStock `xml:"stock"`
}

type SeekingStock struct {
	Symbol  string `xml:"symbol"`
	Company string `xml:"company_name"`
}
