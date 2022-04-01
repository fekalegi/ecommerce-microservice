package domain

type Role struct {
	ID    int    `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}
