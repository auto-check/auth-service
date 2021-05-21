package model

type Student struct {
	ID int32 `json:"id"`
	Gcn string `json:"gcn"`
	Name string `json:"name"`
	Email string `json:"email"`
}