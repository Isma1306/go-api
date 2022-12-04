package types

type Article struct {
	Id      string `json:"_id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
