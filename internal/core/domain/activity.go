package domain

type Activity struct {
	ID         int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title      string  `json:"title"`
	OwnerId    int64   `json:"owner_id"`
	WordSearch [][]int `json:"word_search" gorm:"type:jsonb"`
}
