package responses

type (
	// Response represents the standard structure for API responses.
	Response struct {
		Status      string      `json:"status,omitempty"`
		StatusCode  int         `json:"statusCode,omitempty"`
		Message     string      `json:"message,omitempty"`
		Response    interface{} `json:"response,omitempty"`
		ServiceCode string      `json:"serviceCode,omitempty"`
	}

	// Pagination represents the structure for paginated responses.
	Pagination struct {
		Data       interface{} `json:"data,omitempty"`
		Page       int         `json:"page,omitempty"`
		Size       int         `json:"size,omitempty"`
		TotalPages int         `json:"totalPages,omitempty"`
		TotalItems int         `json:"totalItems,omitempty"`
	}
)
