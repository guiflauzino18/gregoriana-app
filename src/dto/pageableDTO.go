package dto

type Pageable[T any] struct {
	Content          []T  `json:"content"`
	PageNumber       int  `json:"pageNumber"`
	PageSize         int  `json:"pageSize"`
	TotalPages       int  `json:"totalPages"`
	TotalElements    int  `json:"totalElements"`
	First            bool `json:"first"`
	Last             bool `json:"last"`
	NumberOfElements int  `json:"numberOfElements"`
	Empty            bool `json:"empty"`
}
