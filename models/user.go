// models/user.go
package models

import "gorm.io/gorm"

// User struct mendefinisikan tabel User di PostgreSQL
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Age      int    `json:"age"`
	Status   string `json:"status"`
}
