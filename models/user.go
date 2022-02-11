package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	GroupId  string `json:"group_id"`
	Group    *Group `pg:"rel:has-one"`
}
