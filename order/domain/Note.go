package domain

type Note struct {
	ID          int    `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     int    `json:"creator"`

	User User `gorm:"foreignKey:Creator"`
}
