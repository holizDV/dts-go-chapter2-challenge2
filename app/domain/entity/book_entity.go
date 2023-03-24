package entity

type Book struct {
	Id          int
	Title       string
	Author      string
	Description string
}

type BookResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type BookCreatedResponse struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type BookDeleteResponse struct {
	Id int `json:"id"`
}

type BookRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=200"`
	Author      string `json:"author" binding:"required,min=3,max=200"`
	Description string `json:"description" binding:"required,min=5,max=300"`
}
