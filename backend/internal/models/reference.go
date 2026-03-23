package models

type Author struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
