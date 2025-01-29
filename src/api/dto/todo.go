package dto

type CreateToDoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed"`
}

type UpdateToDoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed"`
}

type ToDoResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserId      int    `json:"userId"`
}

type PaginationInputWithFilter struct {
	PageNumber int                    `json:"pageNumber"`
	PageSize   int                    `json:"pageSize"`
	Filter     map[string]interface{} `json:"filter"`
	Sort       map[string]string      `json:"sort"`
}

type PagedList[T any] struct {
	PageNumber  int   `json:"pageNumber"`
	PageSize    int   `json:"pageSize"`
	TotalRows   int64 `json:"totalRows"`
	TotalPages  int   `json:"totalPages"`
	HasNextPage bool  `json:"hasNextPage"`
	HasPrevPage bool  `json:"hasPrevPage"`
	Items       []T   `json:"items"`
}
