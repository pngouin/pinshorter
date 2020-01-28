package models

type Link struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Count    int    `json:"count"`
	ApiPoint string `json:"api_point"`
	User     UserInfo

	Shared
}
