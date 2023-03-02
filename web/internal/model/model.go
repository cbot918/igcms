package model

type Users struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Profile struct {
	Photo string `json:"photo"`
}

type Post struct {
	Title string `json:"title"`
	Content string `json:"content"`
}