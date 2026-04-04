package v1

type Pagination struct {
	Page     int `json:"page" dc:"Current page number"`
	PageSize int `json:"page_size" dc:"Page size"`
	Total    int `json:"total" dc:"Total records"`
}
