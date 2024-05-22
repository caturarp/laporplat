package dto

type Response struct {
	PageCount int    `json:"page_count,omitempty"`
	ItemCount int    `json:"item_count,omitempty"`
	Message   string `json:"message,omitempty"`
	Data      any    `json:"data,omitempty"`
}
type RequestContext struct {
	UserID uint
}

type IDResponse struct {
	ID uint `json:"id"`
}

func NewDataPaginationResponse(data any, pageCount int, itemCount int) (resp *Response) {
	return &Response{PageCount: pageCount, ItemCount: itemCount, Data: data}
}
