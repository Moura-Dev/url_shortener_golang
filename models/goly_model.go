package models

type Goly struct {
	ID          uint64 `json:"id" db:"id"`
	RedirectURL string `json:"redirect_url" db:"redirect_url"`
	Goly        string `json:"goly" db:"goly"`
	Random      bool   `json:"random" db:"random"`
	Count       uint64 `json:"count" db:"count"`
}
