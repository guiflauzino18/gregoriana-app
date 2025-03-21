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

// "last": true,
// 	"totalPages": 1,
// 	"totalElements": 1,
// 	"first": true,
// 	"size": 1,
// 	"number": 0,
// 	"sort": {
// 		"sorted": true,
// 		"empty": false,
// 		"unsorted": false
// 	},
// 	"numberOfElements": 1,
// 	"empty": false
