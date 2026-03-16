package models

type User struct {
	ID    uint   `gorm:"primaryKey;column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

func (User) TableName() string {
	return "users" // ensures GORM queries the correct table
}
