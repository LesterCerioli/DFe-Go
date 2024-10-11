package models

import (
	"github.com/google/uuid"
)

type Content struct {
	ID     uuid.UUID `json:"id"` // UUID instead of int
	Html   string    `json:"html"`
	PageID int       `json:"page_id"`
	Type   int       `json:"type"`
}

type IntranetContentServicePageContent struct {
	ContentID uuid.UUID
	Html      string
	PageID    int
}

func ConvertIntranetContent(item IntranetContentServicePageContent) Content {
	return Content{
		ID:     item.ContentID,
		Html:   item.Html,
		PageID: item.PageID,
	}
}
