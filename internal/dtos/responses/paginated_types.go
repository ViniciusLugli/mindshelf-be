package responses

type PaginatedTaskResponse struct {
	Data        []TaskResponse `json:"data"`
	Total       int64          `json:"total"`
	Page        int            `json:"page"`
	Limit       int            `json:"limit"`
	Total_pages int            `json:"total_pages"`
}

type PaginatedUserResponse struct {
	Data        []UserResponse `json:"data"`
	Total       int64          `json:"total"`
	Page        int            `json:"page"`
	Limit       int            `json:"limit"`
	Total_pages int            `json:"total_pages"`
}

type PaginatedGroupResponse struct {
	Data        []GroupResponse `json:"data"`
	Total       int64           `json:"total"`
	Page        int             `json:"page"`
	Limit       int             `json:"limit"`
	Total_pages int             `json:"total_pages"`
}

type PaginatedMessageResponse struct {
	Data        []MessageResponse `json:"data"`
	Total       int64             `json:"total"`
	Page        int               `json:"page"`
	Limit       int               `json:"limit"`
	Total_pages int               `json:"total_pages"`
}
