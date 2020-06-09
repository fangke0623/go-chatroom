package query

type Page struct {
	PageSize    int `json:"pageSize"`
	CurrentPage int `json:"currentPage"`
}

func GetStartRow(page Page) int {
	return page.PageSize + (page.CurrentPage - 1)
}
