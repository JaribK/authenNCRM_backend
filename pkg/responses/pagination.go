package responses

func ApplyPagination(data interface{}, page, size int) Pagination {
	defaultPage := 1
	defaultSize := 10

	if page <= 0 {
		page = defaultPage
	}
	if size <= 0 {
		size = defaultSize
	}

	items, ok := data.([]interface{})
	if !ok {
		return Pagination{
			Data:       data,
			Page:       page,
			Size:       size,
			TotalPages: 1,
			TotalItems: 1,
		}
	}

	totalItems := len(items)

	start := (page - 1) * size
	end := start + size

	if start > totalItems {
		start = totalItems
	}
	if end > totalItems {
		end = totalItems
	}

	paginatedData := items[start:end]

	totalPages := (totalItems + size - 1) / size

	return Pagination{
		Data:       paginatedData,
		Page:       page,
		Size:       size,
		TotalPages: totalPages,
		TotalItems: totalItems,
	}
}
