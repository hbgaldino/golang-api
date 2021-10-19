package dto

type BookDTO struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
