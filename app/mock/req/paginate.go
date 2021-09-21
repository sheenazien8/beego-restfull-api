package req

type Paginate struct {
	Page    int    `json:"page"`
	Last    int    `json:"last"`
	Limit   int    `json:"limit"`
	Total   int    `json:"total"`
	Keyword string `json:"-"`
}
