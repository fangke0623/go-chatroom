package query

import "strconv"

type Page struct {
	PageSize    string `json:"pageSize"`
	Total       int    `json:"total"`
	TotalPage   int    `json:"totalPage"`
	CurrentPage string `json:"currentPage"`
}

func GetStartRow(page Page) string {
	pageSize, _ := strconv.Atoi(page.PageSize)
	currentPage, _ := strconv.Atoi(page.CurrentPage)
	return strconv.Itoa(pageSize * (currentPage - 1))
}

func SetTotalPage(page Page) Page {
	pageSize, _ := strconv.Atoi(page.PageSize)
	page.TotalPage = page.Total/pageSize + 1
	return page
}
