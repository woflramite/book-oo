package main

type Member struct {
	Students_id string `json:"students_id"`
	Name string `json:"name"`
	Social_media_contact string `json:"social_media_contact"`
}

type Members []Member

type Owns struct {
	Students_id string `json:"students_id"`
	Isbn string `json:"isbn"`
	Number_of_copy string `json:"number_of_copy"`
}

type Ownses []Owns

type Book struct {
	Isbn string `json:"isbn"`
	Book_title string `json:"book_title"`
}

type Books []Book

type BookOwned struct {
	Isbn string `json:"isbn"`
	Book_title string `json:"book_title"`
	Number_of_copy int `json:"number_of_copy"`
}

type BookOwneds []BookOwned