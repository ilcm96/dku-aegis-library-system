package model

type BookReq struct {
	UserId    int
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Reason    string `json:"reason"`
}
