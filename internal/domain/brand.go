package domain

type Brand struct {
	ID   uint `json:"id" gorm:"primaryKey; not null"`
	Name string `json:"name" gorm:"unique not null"`
}
