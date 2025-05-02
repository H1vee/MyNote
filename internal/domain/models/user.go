package models

type User struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Email         string `json:"email" gorm:"not null"`
	Hash_Password string `json:"hash_password" gorm:"not null"`
}
