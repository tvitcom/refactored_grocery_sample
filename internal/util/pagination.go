package util

import "math"

type Pagination struct {
	Model      string
	Count      int
	Page       int
	TotalPages int
	NextPage   int
	PrevPage   int
}

func GetTotalPagesSize(count int, limit int) int {
	pages := float64(count) / float64(limit)
	return int(math.Ceil(pages))
}

func ProcessPagination(model string, count int, page int, limit int) Pagination {
	var pagination Pagination

	pagination.Model = model
	pagination.Count = count
	pagination.Page = page
	pagination.TotalPages = GetTotalPagesSize(count, limit)
	pagination.NextPage = page + 1
	pagination.PrevPage = page - 1

	return pagination
}
