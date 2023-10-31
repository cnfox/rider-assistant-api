package article

import "snail/internal/data"

type listResult struct {
	Category int
	Timeline int
	Search   string
	List     []data.Article
	Page     int
	Total    float64
	Count    int64
}
