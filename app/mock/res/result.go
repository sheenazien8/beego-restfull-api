package res

type Result struct {
	Status   int         `json:"status,omitempty"`
	Code     string      `json:"code,omitempty"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Paginate interface{} `json:"paginate,omitempty"`
}
