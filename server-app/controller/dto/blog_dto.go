package dto

type BlogPostInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type BlogInfo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}