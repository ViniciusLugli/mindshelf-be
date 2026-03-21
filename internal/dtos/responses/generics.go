package responses

type PaginatedResponse[T any] struct {
	Data        []T   `json:"data"`
	Total       int64 `json:"total"`
	Page        int   `json:"page"`
	Limit       int   `json:"limit"`
	Total_pages int   `json:"total_pages"`
}

func NewPaginatedResponse[M any, R any](
	items []M,
	mapper func(M) R,
	total int64,
	page, limit, totalPages int,
) PaginatedResponse[R] {
	data := make([]R, len(items))
	for i, item := range items {
		data[i] = mapper(item)
	}

	return PaginatedResponse[R]{
		Data:        data,
		Total:       total,
		Page:        page,
		Limit:       limit,
		Total_pages: totalPages,
	}
}
