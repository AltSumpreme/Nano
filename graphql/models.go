package main

type Account struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Posts []Post `json:"posts"`
}
